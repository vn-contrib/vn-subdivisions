package resolver

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/vn-contrib/vn-subdivisions/model"
)

type areaResolver struct {
	model.Area
}

func (r areaResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(r.Area.ID, 10))
}

func (r areaResolver) Level() int32 {
	return int32(r.Area.Level)
}
