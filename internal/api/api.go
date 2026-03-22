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

// API implements /pkg/api ServerInterface
type API struct {
	users.UsersAPI
	messages.MessagesAPI

	Ctx    context.Context
	App    App
	Logger *slog.Logger
}

func (api *API) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
	})
}
