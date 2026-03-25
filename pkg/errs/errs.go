// Package errs provides canonical error sentinels, error tagging, HTTP
// mapping, and structured logging/response data utilities.
//
// # Overview
//
// Errors in this package are composable wrappers. A single error value can
// carry a sentinel for routing (via [Mark]), a user-safe message for API
// responses (via [SafeMsg]), extra JSON fields for the response body (via
// [NewRespData]), and structured slog fields for logging (via [NewLogData]).
// All wrappers preserve the original error through the standard Unwrap chain,
// so errors.Is and errors.As continue to work normally.
//
// Typical usage at the infrastructure layer:
//
//	return errs.SafeMsg(errs.Mark(pgErr, errs.ErrExists), "username already exists")
//
// Typical usage at the handler layer:
//
//	if errors.Is(err, errs.ErrExists) {
//	    // errs.GetSafeMsg returns the safe message set above
//	}
//
// # Sentinels
//
// A fixed set of domain-neutral sentinel errors covers the common failure
// categories. Use errors.Is to test for them. The sentinel alone carries no
// user-facing message — attach one with [SafeMsg] when needed.
//
//	ErrNotFound        → 404
//	ErrExists          → 409
//	ErrOutdated        → 409
//	ErrInvalidArgument → 400
//	ErrMissingArgument → 400
//	ErrOutOfRange      → 400
//	ErrUnauthorized    → 401
//	ErrPermissionDenied → 403
//	ErrRateLimited     → 429
//	ErrRemoteServiceErr → 502
//	ErrNotImplemented  → 501
//
// # Marking
//
// [Mark] attaches one or more sentinels to an existing error without changing
// its message or losing the original error in the Unwrap chain:
//
//	err = errs.Mark(err, errs.ErrNotFound)
//	errors.Is(err, errs.ErrNotFound) // true
//
// Multiple sentinels can be applied in one call:
//
//	err = errs.Mark(err, errs.ErrInvalidArgument, errs.ErrMissingArgument)
//
// # Safe messages
//
// Internal errors often contain implementation details (SQL state, stack
// paths, internal IDs) that must not reach API consumers. [SafeMsg] attaches
// a separate user-facing string to any error; [GetSafeMsg] retrieves it.
// [GetHTTPMessageAndStatus] uses the safe message automatically when present,
// falling back to the raw error only for status codes below 500.
//
//	err = errs.SafeMsg(errs.Mark(err, errs.ErrExists), "username already exists")
//	errs.GetSafeMsg(err) // "username already exists"
//	err.Error()          // original internal message (never sent to client)
//
// # HTTP mapping
//
// [GetHTTPCode] maps a sentinel to its HTTP status code.
// [GetHTTPMessageAndStatus] resolves the message using the safe-message
// precedence rules described above. [GetHTTPRespData] returns a ready-to-
// encode JSON response map. [HandleHTTP] does everything — resolves the
// response, logs at the appropriate level, and writes the JSON body.
//
// # Response data
//
// [NewRespData] attaches extra key-value pairs to include in the JSON response
// body alongside the "error" field. Wrappers are immutable — each call
// returns a new error with a merged copy of all accumulated data, with the
// outermost call taking priority on key conflicts:
//
//	err = errs.NewRespData(err, map[string]any{"field": "username"})
//
// # Log data
//
// [NewLogData] attaches slog key-value arguments to an error so they are
// emitted automatically by [Log] and [HandleHTTP]. Wrappers are immutable —
// each call returns a new error with the new args prepended to any previously
// attached args:
//
//	err = errs.NewLogData(err, []any{"user_id", id, "op", "create"})
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

// IsAny reports whether err matches any of the given reference errors via errors.Is.
func IsAny(err error, references ...error) bool {
	for _, reference := range references {
		if errors.Is(err, reference) {
			return true
		}
	}
	return false
}
