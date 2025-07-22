package resolver

import (
	"context"

	"github.com/vn-contrib/vn-subdivisions/gql/ctx"
)

type queryResolver struct{}

func (r *queryResolver) Ping(c context.Context) string {
	db := c.(ctx.Context).DB()
	if err := db.PingContext(c); err != nil {
		panic(err)
	}
	return "pong"
}
