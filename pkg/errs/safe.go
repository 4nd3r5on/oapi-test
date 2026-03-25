package errs

import "errors"

type safeError struct {
	err     error
	safeMsg string
}

func (e *safeError) Error() string {
	return e.err.Error()
}

func (e *safeError) Unwrap() error {
	return e.err
}

// SafeMsg attaches a user-facing message to err. The original error message is
// preserved and still returned by err.Error(); the safe message is only
// accessible via [GetSafeMsg] or [LookupSafeMsg].
//
// Use this to prevent internal details (SQL state, file paths, internal IDs)
// from reaching API consumers while keeping full context for logging.
//
//	err = errs.SafeMsg(errs.Mark(err, errs.ErrExists), "username already exists")
//	errs.GetSafeMsg(err) // "username already exists"
//	err.Error()          // original internal message
func SafeMsg(err error, msg string) error {
	return &safeError{
		err:     err,
		safeMsg: msg,
	}
}

// GetSafeMsg returns the safe message attached to err by [SafeMsg], or an
// empty string if none is present.
func GetSafeMsg(err error) (msg string) {
	if e, ok := errors.AsType[*safeError](err); ok {
		return e.safeMsg
	}
	return ""
}

// LookupSafeMsg is like [GetSafeMsg] but also reports whether a safe message
// was found, allowing callers to distinguish between an empty message and no
// message at all.
func LookupSafeMsg(err error) (msg string, ok bool) {
	if e, ok := errors.AsType[*safeError](err); ok {
		return e.safeMsg, true
	}
	return "", false
}
