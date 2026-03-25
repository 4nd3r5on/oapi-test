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
	UserExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
	UsernameExists(ctx context.Context, username string) (bool, error)
}

type ServiceInterface interface {
	// Upset creates user if doesn't exist and returns the user data
	// or just returns the user
	Upsert(ctx context.Context) (*PrivateUser, error)
	Create(ctx context.Context, opts NewUserOpts) (*PrivateUser, error)

	GetByID(ctx context.Context, id uuid.UUID) (*PublicUser, error)
	GetByUsername(ctx context.Context, username string) (*PublicUser, error)

	GetMe(ctx context.Context) (*PrivateUser, error)
	UpdateMe(ctx context.Context, opts UpdateUserOpts) (*PrivateUser, error)
	DeleteMe(ctx context.Context) error

	IsUsernameAvailable(ctx context.Context, username string) (bool, error)
}
