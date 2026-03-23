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

func SafeMsg(err error, msg string) error {
	return &safeError{
		err:     err,
		safeMsg: msg,
	}
}

func GetSafeMsg(err error) (msg string) {
	if e, ok := errors.AsType[*safeError](err); ok {
		return e.safeMsg
	}
	return ""
}

func LookupSafeMsg(err error) (msg string, ok bool) {
	if e, ok := errors.AsType[*safeError](err); ok {
		return e.safeMsg, true
	}
	return "", false
}
