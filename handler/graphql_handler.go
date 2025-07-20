package handler

import (
	"encoding/json"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

var _ http.Handler = (*GraphQLHandler)(nil)

type GraphQLHandler struct {
	Schema *graphql.Schema
}

func (h *GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
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

		res := h.Schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(res)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
