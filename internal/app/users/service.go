// Package users provides application layer for users
package users

import (
	"context"

	"github.com/google/uuid"

	"github.com/4nd3rs0n/oapi-test/internal/app/auth"
	"github.com/4nd3rs0n/oapi-test/pkg/errs"
)

type Service struct {
	db DB
}

func NewService(db DB) *Service {
	return &Service{db: db}
}

// resolveUserID resolves the caller's user ID from ClientData stored in ctx.
// For TMA: upserts user from Telegram init data.
// For TMB: verifies the claimed user ID exists in the DB.
// For Session: uses the user ID stored in the session.
func (s *Service) resolveUserID(ctx context.Context) (uuid.UUID, error) {
	clientData, ok := auth.CtxGetClientData(ctx)
	if !ok {
		return uuid.Nil, errs.ErrUnauthorized
	}
	switch clientData.Method {
	case auth.MethodTMA:
		// TODO: upsert user via TG init data, return user ID
		panic("Not implemented")
	case auth.MethodTMB:
		id := clientData.TMBData.UserID
		exists, err := s.db.UserExistsByID(ctx, id)
		if err != nil {
			return uuid.Nil, err
		}
		if !exists {
			return uuid.Nil, errs.ErrUnauthorized
		}
		return id, nil
	case auth.MethodSession:
		panic("Not implemented")
	default:
		return uuid.Nil, errs.ErrUnauthorized
	}
}

// Upsert creates user if doesn't exist and returns the user data, or just
// returns the existing user. Only valid for TMA auth.
func (s *Service) Upsert(ctx context.Context) (*PrivateUser, error) {
	clientData, ok := auth.CtxGetClientData(ctx)
	if !ok {
		return nil, errs.ErrUnauthorized
	}
	if clientData.Method != auth.MethodTMA {
		return nil, errs.Mark(errs.ErrPermissionDenied)
	}
	// TODO: upsert user via TG init data
	panic("Not implemented")
}

func (s *Service) Create(ctx context.Context, opts NewUserOpts) (*PrivateUser, error) {
	// TODO: Change how username is validated
	if err := validateUsername(opts.Username); err != nil {
		return nil, err
	}
	if err := validateBio(opts.Bio); err != nil {
		return nil, err
	}
	user, err := s.db.CreateUser(ctx, opts)
	if err != nil {
		return nil, err
	}
	return user.PrivView(), nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*PublicUser, error) {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user.PubView(), nil
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*PublicUser, error) {
	user, err := s.db.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user.PubView(), nil
}

func (s *Service) GetMe(ctx context.Context) (*PrivateUser, error) {
	id, err := s.resolveUserID(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user.PrivView(), nil
}

func (s *Service) UpdateMe(ctx context.Context, opts UpdateUserOpts) (*PrivateUser, error) {
	id, err := s.resolveUserID(ctx)
	if err != nil {
		return nil, err
	}
	if opts.IsEmpty() {
		return nil, errs.SafeMsg(ErrNoFieldsToUpdate, ErrNoFieldsToUpdate.Error())
	}
	if opts.Username != nil {
		if err = validateUsername(*opts.Username); err != nil {
			return nil, err
		}
	}
	if opts.Bio != nil {
		if err = validateBio(*opts.Bio); err != nil {
			return nil, err
		}
	}
	user, err := s.db.UpdateUser(ctx, id, opts)
	if err != nil {
		return nil, err
	}
	return user.PrivView(), nil
}

func (s *Service) DeleteMe(ctx context.Context) error {
	id, err := s.resolveUserID(ctx)
	if err != nil {
		return err
	}
	return s.db.DeleteUser(ctx, id)
}

// IsUsernameAvailable reports whether the given username is not taken.
// Returns ErrInvalidUsername if the username violates length constraints.
func (s *Service) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	if len(username) < UsernameMinLen || len(username) > UsernameMaxLen {
		return false, ErrInvalidUsername
	}
	exists, err := s.db.UsernameExists(ctx, username)
	if err != nil {
		return false, err
	}
	return !exists, nil
}
