package auth

import (
	"context"
	"strings"

	"github.com/4nd3rs0n/oapi-test/pkg/errs"
	"github.com/google/uuid"
)

// TMB stands for Trust Me Bro auth
type TMB struct{}

type TMBData struct {
	UserID uuid.UUID
}

func (auth *TMB) Verify(ctx context.Context, _, token string, _ []string) (*ClientData, error) {
	val, ok := strings.CutPrefix(token, "TMB ")
	if !ok {
		return nil, errs.Mark(errs.ErrInvalidArgument)
	}
	userID, err := uuid.Parse(val)
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
