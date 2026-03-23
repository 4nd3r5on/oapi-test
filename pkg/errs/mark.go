package errs

import "errors"

// markedError wraps a base error with one or more sentinels for matching.
type markedError struct {
	err       error
	sentinels []error
}

// Error returns the original error message, satisfying the error interface.
func (m *markedError) Error() string {
	return m.err.Error()
}

// Unwrap returns the underlying error so errors.As and other wrapping logic works.
func (m *markedError) Unwrap() error {
	return m.err
}

// Is checks if the target matches any of the assigned sentinels.
func (m *markedError) Is(target error) bool {
	for _, s := range m.sentinels {
		if errors.Is(s, target) || s == target {
			return true
		}
	}
	return false
}

// Mark marks an error with one or more sentinel errors for errors.Is matching.
// Returns nil if the first error (the subject) is nil.
// Usage: Mark(originalErr, sentinelOne, sentinelTwo)
func Mark(errs ...error) error {
	if len(errs) == 0 || errs[0] == nil {
		return nil
	}

	// If no sentinels are provided, return the error as is.
	if len(errs) == 1 {
		return errs[0]
	}

	return &markedError{
		err:       errs[0],
		sentinels: errs[1:],
	}
}
