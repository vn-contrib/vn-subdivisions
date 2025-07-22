package gql

import (
	_ "embed"

	"github.com/graph-gophers/graphql-go"
	"github.com/vn-contrib/vn-subdivisions/gql/resolver"
)

//go:embed schema.gql
var schemaString string

func newSchema() *graphql.Schema {
	return graphql.MustParseSchema(schemaString, resolver.NewRootResolver(),
		graphql.UseFieldResolvers(),
	)
}
