package api

import (
	"github.com/4nd3rs0n/oapi-test/internal/api/messages"
	"github.com/4nd3rs0n/oapi-test/internal/api/users"
)

type App interface {
	users.App
	messages.App
}
