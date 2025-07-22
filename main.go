package main

import (
	"net/http"
	"os"

	_ "github.com/errybase/go-dotenv/autoload"
	"github.com/vn-contrib/vn-subdivisions/db"
	"github.com/vn-contrib/vn-subdivisions/gql"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	gqlHdr := gql.NewHandler(db)
	http.Handle("/graphql", gqlHdr)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
