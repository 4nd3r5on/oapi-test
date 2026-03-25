package auth

import (
	"context"

	"github.com/4nd3rs0n/oapi-test/internal/app/sessions"
	"github.com/4nd3rs0n/oapi-test/pkg/errs"
)

type Sessions struct{}

func (s *Sessions) Verify(ctx context.Context, _ string, sessionKey string, _ []string) (*ClientData, error) {
	isValid := sessions.ValidateKey(sessionKey)
	if !isValid {
		return nil, errs.SafeMsg(errs.Mark(errs.ErrUnauthorized), "invalid or expired session")
	}
	return &ClientData{
		Method:      MethodSession,
		SessionData: &SessionData{Key: sessionKey},
	}, nil
}
