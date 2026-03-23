package errs

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
)

func GetHTTPCode(err error) int {
	switch {
	case errors.Is(err, ErrNotImplemented):
		return http.StatusNotImplemented
	case errors.Is(err, context.DeadlineExceeded):
		return http.StatusGatewayTimeout
	case errors.Is(err, ErrRemoteServiceErr):
		return http.StatusBadGateway
	case errors.Is(err, ErrRateLimited):
		return http.StatusTooManyRequests
	case IsAny(err,
		ErrInvalidArgument,
		ErrMissingArgument,
		ErrOutOfRange,
	):
		return http.StatusBadRequest
	case errors.Is(err, ErrPermissionDenied):
		return http.StatusForbidden
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case IsAny(err, ErrExists, ErrOutdated):
		return http.StatusConflict
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func IsSafeCode(stats int) bool {
	return stats < 500
}

func HTTPGetLogLevel(status int) slog.Level {
	switch {
	case status >= 500:
		return slog.LevelError
	case status == http.StatusUnauthorized || status == http.StatusForbidden:
		return slog.LevelWarn
	default:
		return slog.LevelDebug
	}
}

func GetHTTPData(err error) (status int, msg string) {
	statusCode := GetHTTPCode(err)
	if safeMsg := GetSafeMsg(err); safeMsg != "" {
		return statusCode, safeMsg
	}
	if !IsSafeCode(statusCode) {
		return statusCode, http.StatusText(statusCode)
	}
	return statusCode, err.Error()
}

func GetHTTPMsg(err error) string {
	_, msg := GetHTTPData(err)
	return msg
}
