package errs

import "errors"

var (
	ErrNotImplemented   = errors.New("not implemented")
	ErrRemoteServiceErr = errors.New("remote service error")
	ErrRateLimited      = errors.New("rate limited")

	ErrInvalidArgument = errors.New("invalid argument")
	ErrMissingArgument = errors.New("missing argument")
	ErrOutOfRange      = errors.New("out of range")

	ErrPermissionDenied = errors.New("permission denied")
	ErrUnauthorized     = errors.New("unauthorized")

	ErrExists   = errors.New("already exists")
	ErrNotFound = errors.New("not found")
	ErrOutdated = errors.New("outdated")
)

func IsAny(err error, references ...error) bool {
	for _, reference := range references {
		if errors.Is(err, reference) {
			return true
		}
	}
	return false
}
