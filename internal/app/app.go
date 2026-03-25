// Package app provides high-level orchestration for the application layer
// aka usecases
package app

import "github.com/4nd3rs0n/oapi-test/internal/app/users"

type DB interface {
	users.DB
}

type Cache interface{}

type Storage interface{}

type App struct {
	Users *users.Service
}

func New(db DB, cache Cache, storage Storage) (*App, error) {
	users := users.NewService(db)
	// sessions.NewService(cache)

	return &App{
		Users: users,
		// TODO:
		// - Posts
		// - Usernames
	}, nil
}
