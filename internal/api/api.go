// Package api provides /pkg/api server implementation
package api

import (
	"context"
	"encoding/json"
	"net/http"
)

type Dependencies struct{}

// API implements /pkg/api ServerInterface
type API struct{}

func NewAPI(ctx context.Context) *API {
	return &API{}
}

func (api *API) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
	})
}
