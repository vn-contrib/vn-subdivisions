package resolver

import (
	"context"

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

func (*queryResolver) Subdivisions(c context.Context) ([]areaResolver, error) {
	db := c.(ctx.Context).DB()
	var resolvers []areaResolver
	var areas []model.Area

	if err := db.NewSelect().Model(&areas).Scan(c); err != nil {
		return resolvers, err
	}

	for _, area := range areas {
		resolvers = append(resolvers, areaResolver{area})
	}

	return resolvers, nil
}
