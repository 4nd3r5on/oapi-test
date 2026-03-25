package api

import (
	"context"

	"github.com/4nd3rs0n/oapi-test/pkg/api"
	"github.com/4nd3rs0n/oapi-test/pkg/errs"
)

func (h *APIHandler) NewError(ctx context.Context, err error) *api.InternalErrorStatusCode {
	status, msg := errs.GetHTTPMessageAndStatus(err)
	logLevel := errs.GetHTTPLogLevel(status)
	errs.Log(ctx, h.Logger, logLevel, err)
	return &api.InternalErrorStatusCode{
		StatusCode: status,
		Response:   api.ErrorResponse{Error: msg},
	}
}
