package users

import "errors"

const (
	UsernameMinLen = 1
	UsernameMaxLen = 64
	BioMaxLen      = 512
)

var (
	ErrNoFieldsToUpdate = errors.New("at least one field must be provided for update")
	ErrInvalidUsername  = errors.New("username must be between 1 and 64 characters")
)

// PubView returns the public representation of the user.
func (u *User) PubView() *PublicUser {
	return &PublicUser{
		ID:       u.ID,
		Username: u.Username,
		Bio:      u.Bio,
	}
}

// PrivView returns the private representation of the user.
func (u *User) PrivView() *PrivateUser {
	return &PrivateUser{
		PublicUser: *u.PubView(),
		Locale:     u.Locale,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

// IsEmpty reports whether no fields are set, which is invalid per the API contract.
func (o UpdateUserOpts) IsEmpty() bool {
	return o.Username == nil && o.Locale == nil && o.Bio == nil
}
