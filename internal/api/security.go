package api

import (
	"context"
	"fmt"

	"github.com/ogen-go/ogen/ogenerrors"

	"github.com/4nd3rs0n/oapi-test/internal/app/auth"
	"github.com/4nd3rs0n/oapi-test/pkg/api"
	pkgerrs "github.com/4nd3rs0n/oapi-test/pkg/errs"
)

type SecurityHandler struct {
	TMA, TMB, Session auth.VerifierFunc
}

func (sh *SecurityHandler) HandleSessionKey(
	ctx context.Context,
	operationName api.OperationName,
	t api.SessionKey,
) (context.Context, error) {
	if sh.Session == nil {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}
	clientData, err := sh.Session(ctx, string(operationName), t.Token, t.Roles)
	return handleVerifierFuncOut(ctx, clientData, err)
}

func (sh *SecurityHandler) HandleTelegramMiniApp(ctx context.Context, operationName api.OperationName, t api.TelegramMiniApp) (context.Context, error) {
	if sh.TMA == nil {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}
	clientData, err := sh.TMA(ctx, string(operationName), t.APIKey, t.Roles)
	return handleVerifierFuncOut(ctx, clientData, err)
}

func (sh *SecurityHandler) HandleTrustMeBro(ctx context.Context, operationName api.OperationName, t api.TrustMeBro) (context.Context, error) {
	if sh.TMB == nil {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}
	clientData, err := sh.TMB(ctx, string(operationName), t.APIKey, t.Roles)
	return handleVerifierFuncOut(ctx, clientData, err)
}

func handleVerifierFuncOut(ctx context.Context, clientData *auth.ClientData, err error) (context.Context, error) {
	if err != nil {
		return ctx, authErr(err)
	}
	return auth.CtxPutClientData(ctx, clientData), nil
}

// authErr maps a verifier error to the appropriate ogen security error.
// Malformed tokens skip the current scheme so ogen can try the next one.
// Denied or internal errors fail the request immediately.
func authErr(err error) error {
	if pkgerrs.IsAny(err, pkgerrs.ErrInvalidArgument, pkgerrs.ErrMissingArgument) {
		return ogenerrors.ErrSkipServerSecurity
	}
	if pkgerrs.IsAny(err, pkgerrs.ErrPermissionDenied, pkgerrs.ErrUnauthorized) {
		return fmt.Errorf("%w", err)
	}
	return fmt.Errorf("auth: %w", err)
}
