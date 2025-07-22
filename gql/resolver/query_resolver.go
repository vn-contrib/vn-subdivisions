package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/vn-contrib/vn-subdivisions/gql/ctx"
	"github.com/vn-contrib/vn-subdivisions/model"
)

type queryResolver struct{}

func (*queryResolver) Ping(c context.Context) string {
	db := c.(ctx.Context).DB()
	if err := db.PingContext(c); err != nil {
		panic(err)
	}
	return "pong"
}

func (*queryResolver) Subdivisions(c context.Context, vars struct {
	Filters *submissionFilters
}) ([]areaResolver, error) {
	db := c.(ctx.Context).DB()
	var resolvers []areaResolver
	var areas []model.Area

	query := db.NewSelect().Model(&areas)
	if filters := vars.Filters; filters != nil {
		if filters.ParentID.Set {
			query = query.Where("parent_id = ?", *filters.ParentID.Value)
		}
		if filters.Level.Set {
			query = query.Where("level = ?", filters.Level.Value)
		}
		if filters.Unit.Set {
			query = query.Where("unit = ?", filters.Unit.Value)
		}
	}

	if err := query.Scan(c); err != nil {
		return resolvers, err
	}

	for _, area := range areas {
		resolvers = append(resolvers, areaResolver{area})
	}

	return resolvers, nil
}

type submissionFilters struct {
	ParentID graphql.NullID
	Level    graphql.NullInt
	Unit     graphql.NullString
}
