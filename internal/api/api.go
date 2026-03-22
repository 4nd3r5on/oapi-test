// Package api provides /pkg/api server implementation
package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/4nd3rs0n/oapi-test/internal/api/messages"
	"github.com/4nd3rs0n/oapi-test/internal/api/users"
)

// APIHandler implements /pkg/api ServerInterface
type APIHandler struct {
	users.UsersAPI
	messages.MessagesAPI

	ctx context.Context
}

func NewAPIHandler(ctx context.Context, app App, logger *slog.Logger) (*APIHandler, error) {
	api := &APIHandler{}
	return api, nil
}

func (api *APIHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
	})
}
