package main

import (
	"net/http"
	"os"

	_ "github.com/errybase/go-dotenv/autoload"
	"github.com/vn-contrib/vn-subdivisions/gql"
	"github.com/vn-contrib/vn-subdivisions/handler"
)

func main() {
	schema := gql.NewSchema()
	http.Handle("/graphql", &handler.GraphQLHandler{Schema: schema})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
