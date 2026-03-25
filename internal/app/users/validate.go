package users

import (
	"github.com/4nd3rs0n/oapi-test/pkg/errs"
)

func validateUsername(username string) error {
	// TODO: Allow only english letters, numbers and _
	if len(username) < UsernameMinLen || len(username) > UsernameMaxLen {
		return errs.SafeMsg(errs.Mark(ErrInvalidUsername, errs.ErrInvalidArgument), ErrInvalidUsername.Error())
	}
	return nil
}

func validateBio(bio string) error {
	if len(bio) > BioMaxLen {
		return errs.SafeMsg(errs.Mark(errs.ErrInvalidArgument), "bio must not exceed 512 characters")
	}
	return nil
}
