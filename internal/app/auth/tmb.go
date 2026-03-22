package auth

import (
	"context"

	"github.com/4nd3r5on/errs"
	"github.com/google/uuid"
)

type AuthTMBRepo interface {
	ExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
}

// TMB stands for Trust Me Bro auth
type TMB struct{}

type TMBData struct {
	UserID uuid.UUID
}

func (auth *TMB) Verify(ctx context.Context, data string) (*ClientData, error) {
	// data is expected to be a raw UUID string
	userID, err := uuid.Parse(data)
	if err != nil {
		return nil, errs.F().
			Message("invalid user id format").
			Mark(errs.ErrInvalidArgument).
			Err()
	}

	return &ClientData{
		Method: MethodTMB,
		TMBData: &TMBData{
			UserID: userID,
		},
	}, nil
}
