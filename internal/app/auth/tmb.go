package auth

import (
	"context"

	"github.com/4nd3rs0n/oapi-test/pkg/errs"
	"github.com/google/uuid"
)

// TMB stands for Trust Me Bro auth
type TMB struct{}

type TMBData struct {
	UserID uuid.UUID
}

func (auth *TMB) Verify(ctx context.Context, _, userIDStr string, _ []string) (*ClientData, error) {
	// data is expected to be a raw UUID string
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		err = errs.Mark(errs.ErrInvalidArgument)
		err = errs.SafeMsg(err, "Trust Me Bro auth invalid user id format")
		return nil, err
	}

	return &ClientData{
		Method: MethodTMB,
		TMBData: &TMBData{
			UserID: userID,
		},
	}, nil
}
