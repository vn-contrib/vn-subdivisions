package fixtures

import (
	"context"
	"database/sql"
	"encoding/csv"
	"os"

	"github.com/uptrace/bun"
	"github.com/vn-contrib/vn-subdivisions/model"
)

func init() {
	Fixtures.Register(func(ctx context.Context, db *bun.DB) error {
		file, err := os.Open("cmd/db/fixtures/static/subdivisions_20250710.csv")
		if err != nil {
			return err
		}
		defer file.Close()

		rows, err := csv.NewReader(file).ReadAll()
		if err != nil {
			return err
		}

		var lv1Areas, lv2Areas []*model.Area

		for i, row := range rows {
			if i == 0 {
				continue
			}

			gsoID, name, unit, parentGsoID := row[0], row[1], row[3], row[4]

			if gsoID == "" || gsoID == parentGsoID {
				continue
			}

			area := &model.Area{
				Name:        name,
				Unit:        unit,
				GsoID:       gsoID,
				ParentGsoID: parentGsoID,
			}

			if parentGsoID == "" {
				area.Level = 1
				lv1Areas = append(lv1Areas, area)
			} else {
				area.Level = 2
				lv2Areas = append(lv2Areas, area)
			}
		}

		return db.RunInTx(context.Background(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
			if _, err := tx.NewInsert().
				Model(&lv1Areas).
				On("CONFLICT (gso_id) DO UPDATE").
				Set("name = EXCLUDED.name, unit = EXCLUDED.unit, level = EXCLUDED.level").
				Exec(ctx); err != nil {
				return err
			}

			lv1AreasByGsoID := map[string]*model.Area{}
			for _, area := range lv1Areas {
				lv1AreasByGsoID[area.GsoID] = area
			}

			for _, area := range lv2Areas {
				area.ParentID = lv1AreasByGsoID[area.ParentGsoID].ID
			}

			if _, err := tx.NewInsert().
				Model(&lv2Areas).
				On("CONFLICT (gso_id) DO UPDATE").
				Set("name = EXCLUDED.name, unit = EXCLUDED.unit, level = EXCLUDED.level, parent_id = EXCLUDED.parent_id").
				Exec(ctx); err != nil {
				return err
			}

			return nil
		})
	})
}
