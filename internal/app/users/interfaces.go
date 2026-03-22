package users

import (
	"context"

	"github.com/google/uuid"
)

type DB interface {
	CreateUser(ctx context.Context, opts NewUserOpts) (*User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, opts UpdateUserOpts) (*User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UsernameExists(ctx context.Context, username string) (bool, error)
}

type ServiceInterface interface {
	Create(ctx context.Context, opts NewUserOpts) (*PrivUser, error)
	GetPrivByID(ctx context.Context, id uuid.UUID) (*PrivUser, error)
	GetPubByID(ctx context.Context, id uuid.UUID) (*PubUser, error)
	GetPubByUsername(ctx context.Context, username string) (*PubUser, error)
	Update(ctx context.Context, id uuid.UUID, opts UpdateUserOpts) (*PrivUser, error)
	Delete(ctx context.Context, id uuid.UUID) error
	IsUsernameAvailable(ctx context.Context, username string) (bool, error)
}
