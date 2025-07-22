package gql

import (
	"encoding/json"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/uptrace/bun"
	"github.com/vn-contrib/vn-subdivisions/gql/ctx"
)

var _ http.Handler = (*Handler)(nil)

func NewHandler(db *bun.DB) *Handler {
	return &Handler{
		schema: newSchema(),
		db:     db,
	}
}

type Handler struct {
	schema *graphql.Schema
	db     *bun.DB
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "static/graphiql.html")
	case http.MethodPost:
		var params struct {
			Query         string         `json:"query"`
			OperationName string         `json:"operationName"`
			Variables     map[string]any `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := ctx.NewContext(r.Context(), h.db)
		res := h.schema.Exec(ctx, params.Query, params.OperationName, params.Variables)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(res)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
