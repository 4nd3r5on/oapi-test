package errs

import (
	"context"
	"errors"
	"log/slog"
)

type logDataError struct {
	err      error
	slogArgs []any
}

func (e *logDataError) Error() string {
	return e.err.Error()
}

func (e *logDataError) Unwrap() error {
	return e.err
}

// NewLogData attaches structured slog key-value arguments to err so they are
// emitted automatically when [Log] or [HandleHTTP] processes it. If err
// already carries a [logDataError], the existing args are appended after the
// new ones.
//
// Each call is immutable — a fresh slice is always allocated, so the original
// error and any previously attached args are never mutated.
//
//	err = errs.NewLogData(err, []any{"user_id", id, "op", "create"})
func NewLogData(err error, slogArgs []any) error {
	merged := make([]any, len(slogArgs))
	copy(merged, slogArgs)
	if ldErr, ok := errors.AsType[*logDataError](err); ok {
		merged = append(merged, ldErr.slogArgs...)
	}
	return &logDataError{
		err:      err,
		slogArgs: merged,
	}
}

// Log emits err to logger at the given level. Any slog args attached via
// [NewLogData] are appended to args automatically.
func Log(ctx context.Context, logger *slog.Logger, level slog.Level, err error, args ...any) {
	if ldErr, ok := errors.AsType[*logDataError](err); ok && ldErr.slogArgs != nil {
		if args == nil {
			args = ldErr.slogArgs
		} else {
			args = append(args, ldErr.slogArgs...)
		}
	}
	logger.Log(ctx, level, err.Error(), args...)
}
