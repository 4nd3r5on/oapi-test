package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	appusers "github.com/4nd3rs0n/oapi-test/internal/app/users"
	dbpkg "github.com/4nd3rs0n/oapi-test/pkg/db"
)

type Users struct {
	q *dbpkg.Queries
}

func (u *Users) CreateUser(ctx context.Context, opts appusers.NewUserOpts) (*appusers.User, error) {
	row, err := u.q.CreateUser(ctx, dbpkg.CreateUserParams{
		ID:       uuidToPg(opts.ID),
		Username: strToPg(opts.Username),
		Locale:   opts.Locale,
		Bio:      opts.Bio,
	})
	if err != nil {
		return nil, fmt.Errorf("create user: %w", mapErr(err))
	}
	return dbUserToApp(row), nil
}

func (u *Users) GetUserByID(ctx context.Context, id uuid.UUID) (*appusers.User, error) {
	row, err := u.q.GetUserByID(ctx, uuidToPg(id))
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", mapErr(err))
	}
	return dbUserToApp(row), nil
}

func (u *Users) GetUserByUsername(ctx context.Context, username string) (*appusers.User, error) {
	row, err := u.q.GetUserByUsername(ctx, strToPg(username))
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", mapErr(err))
	}
	return dbUserToApp(row), nil
}

func (u *Users) UpdateUser(ctx context.Context, id uuid.UUID, opts appusers.UpdateUserOpts) (*appusers.User, error) {
	row, err := u.q.UpdateUser(ctx, dbpkg.UpdateUserParams{
		ID:       uuidToPg(id),
		Username: ptrStrToPg(opts.Username),
		Locale:   ptrStrToPg(opts.Locale),
		Bio:      ptrStrToPg(opts.Bio),
	})
	if err != nil {
		return nil, fmt.Errorf("update user: %w", mapErr(err))
	}
	return dbUserToApp(row), nil
}

func (u *Users) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := u.q.DeleteUser(ctx, uuidToPg(id)); err != nil {
		return fmt.Errorf("delete user: %w", mapErr(err))
	}
	return nil
}

func (u *Users) UserExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	exists, err := u.q.UserExistsByID(ctx, uuidToPg(id))
	if err != nil {
		return false, fmt.Errorf("user exists by id: %w", mapErr(err))
	}
	return exists, nil
}

func (u *Users) UsernameExists(ctx context.Context, username string) (bool, error) {
	exists, err := u.q.UsernameExists(ctx, strToPg(username))
	if err != nil {
		return false, fmt.Errorf("username exists: %w", mapErr(err))
	}
	return exists, nil
}
