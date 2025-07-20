package main

import (
	"net/http"
	"os"

	_ "github.com/errybase/go-dotenv/autoload"
)

func main() {
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
