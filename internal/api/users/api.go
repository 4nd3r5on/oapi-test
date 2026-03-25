// Package users provides users API
package users

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/4nd3rs0n/oapi-test/pkg/errs"
	"github.com/google/uuid"

	appusers "github.com/4nd3rs0n/oapi-test/internal/app/users"
	"github.com/4nd3rs0n/oapi-test/pkg/api"
)

type UsersAPI struct {
	Users  appusers.ServiceInterface
	Logger *slog.Logger
}

func (u *UsersAPI) CheckUsernameAvailability(ctx context.Context, params api.CheckUsernameAvailabilityParams) (api.CheckUsernameAvailabilityRes, error) {
	available, err := u.Users.IsUsernameAvailable(ctx, params.Username)
	if err != nil {
		if errors.Is(err, appusers.ErrInvalidUsername) {
			return newErrorResponse(err), nil
		}
		return nil, fmt.Errorf("check username availability: %w", err)
	}
	return &api.UsernameAvailability{Available: available}, nil
}

func (u *UsersAPI) CreateUser(ctx context.Context, req *api.NewUserRequest) (api.CreateUserRes, error) {
	opts := appusers.NewUserOpts{
		ID:       uuid.UUID(req.ID),
		Username: req.Username.Or(""),
		Locale:   req.Locale,
		Bio:      req.Bio.Or(""),
	}
	user, err := u.Users.Create(ctx, opts)
	if err != nil {
		if errors.Is(err, errs.ErrExists) {
			c := api.NewResourceConflictErrorConflict(api.ResourceConflictError{
				Code:  api.OptResourceConflictErrorCode{Set: true, Value: api.ResourceConflictErrorCodeResourceExists},
				Error: errs.GetSafeMsg(err),
			})
			return &c, nil
		}
		if errs.IsAny(err, errs.ErrInvalidArgument, errs.ErrMissingArgument) {
			return newErrorResponse(err), nil
		}
		return nil, fmt.Errorf("create user: %w", err)
	}
	return privateUserToAPI(user), nil
}

func (u *UsersAPI) DeleteMe(ctx context.Context) error {
	if err := u.Users.DeleteMe(ctx); err != nil {
		return fmt.Errorf("delete me: %w", err)
	}
	return nil
}

func (u *UsersAPI) GetMe(ctx context.Context) (*api.PrivUser, error) {
	user, err := u.Users.GetMe(ctx)
	if err != nil {
		return nil, fmt.Errorf("get me: %w", err)
	}
	return privateUserToAPI(user), nil
}

func (u *UsersAPI) GetUserByID(ctx context.Context, params api.GetUserByIDParams) (api.GetUserByIDRes, error) {
	user, err := u.Users.GetByID(ctx, params.UserID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return newErrorResponse(err), nil
		}
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return publicUserToAPI(user), nil
}

func (u *UsersAPI) GetUserByUsername(ctx context.Context, params api.GetUserByUsernameParams) (api.GetUserByUsernameRes, error) {
	user, err := u.Users.GetByUsername(ctx, params.Username)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return newErrorResponse(err), nil
		}
		return nil, fmt.Errorf("get user by username: %w", err)
	}
	return publicUserToAPI(user), nil
}

func (u *UsersAPI) UpdateMe(ctx context.Context, req *api.UpdateUserRequest) (api.UpdateMeRes, error) {
	opts := appusers.UpdateUserOpts{}
	if v, ok := req.Username.Get(); ok {
		opts.Username = &v
	}
	if v, ok := req.Locale.Get(); ok {
		opts.Locale = &v
	}
	if v, ok := req.Bio.Get(); ok {
		opts.Bio = &v
	}
	user, err := u.Users.UpdateMe(ctx, opts)
	if err != nil {
		if errors.Is(err, errs.ErrExists) {
			c := api.NewResourceConflictErrorConflict(api.ResourceConflictError{
				Code:  api.OptResourceConflictErrorCode{Set: true, Value: api.ResourceConflictErrorCodeResourceExists},
				Error: errs.GetSafeMsg(err),
			})
			return &c, nil
		}
		if errs.IsAny(err, errs.ErrInvalidArgument, errs.ErrMissingArgument) || errors.Is(err, appusers.ErrNoFieldsToUpdate) {
			return newErrorResponse(err), nil
		}
		return nil, fmt.Errorf("update me: %w", err)
	}
	return privateUserToAPI(user), nil
}

func newErrorResponse(err error) *api.ErrorResponse {
	return &api.ErrorResponse{Error: errs.GetSafeMsg(err)}
}
