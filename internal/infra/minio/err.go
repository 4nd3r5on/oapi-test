package minio

import (
	"context"
	"errors"
	"net/http"

	"github.com/4nd3rs0n/oapi-test/pkg/errs"
	"github.com/minio/minio-go/v7"
)

func mapErr(err error) error {
	if err == nil {
		return nil
	}

	errResp, ok := errors.AsType[*minio.ErrorResponse](err)
	if ok {
		switch errResp.StatusCode {
		case http.StatusNotFound:
			return errs.Mark(errResp, errs.ErrNotFound)
		case http.StatusForbidden:
			return errs.Mark(errResp, errs.ErrPermissionDenied)
		case http.StatusUnauthorized:
			return errs.Mark(errResp, errs.ErrUnauthorized)
		case http.StatusTooManyRequests:
			return errs.Mark(errResp, errs.ErrRateLimited)
		case http.StatusConflict:
			return errs.Mark(errResp, errs.ErrExists)
		default:
			return errs.Mark(errResp, errs.ErrRemoteServiceErr)
		}
	}

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}

	return errs.Mark(err, errs.ErrRemoteServiceErr)
}
