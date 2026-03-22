// Package app provides high-level orchestration for the application layer
// aka usecases
package app

type App struct{}

func New() *App {
	return &App{}
}
