package fixtures

import (
	"context"

	"github.com/uptrace/bun"
)

type Seeder struct {
	DB       *bun.DB
	Fixtures *fixtures
}

func (s *Seeder) Seed(ctx context.Context) error {
	for _, seed := range *s.Fixtures {
		if err := seed(ctx, s.DB); err != nil {
			return err
		}
	}
	return nil
}
