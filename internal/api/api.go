// Package api provides /pkg/api server implementation
package api

import (
	"log/slog"

	"github.com/4nd3rs0n/oapi-test/internal/api/users"
	"github.com/4nd3rs0n/oapi-test/internal/app"
)

// APIHandler implements /pkg/api ServerInterface
type APIHandler struct {
	*users.UsersAPI

	logger *slog.Logger
}

func NewAPIHandler(appLayer *app.App, logger *slog.Logger) (*APIHandler, error) {
	usersAPI := &users.UsersAPI{
		Users:  appLayer.Users,
		Logger: logger,
	}

	api := &APIHandler{
		UsersAPI: usersAPI,
		logger:   logger,
	}
	return api, nil
}
