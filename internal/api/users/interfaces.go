package users

import "context"

type App interface {
	CreateUser(ctx context.Context) error
	GetMe(ctx context.Context) error
	GetUser(ctx context.Context) error
}
