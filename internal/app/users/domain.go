package users

import "errors"

const (
	UsernameMinLen = 1
	UsernameMaxLen = 64
	BioMaxLen      = 512
)

var ErrNoFieldsToUpdate = errors.New("at least one field must be provided for update")

// PubView returns the public representation of the user.
func (u *User) PubView() *PubUser {
	return &PubUser{
		ID:       u.ID,
		Username: u.Username,
		Bio:      u.Bio,
	}
}

// PrivView returns the private representation of the user.
func (u *User) PrivView() *PrivUser {
	return &PrivUser{
		PubUser:   *u.PubView(),
		Locale:    u.Locale,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// IsEmpty reports whether no fields are set, which is invalid per the API contract.
func (o UpdateUserOpts) IsEmpty() bool {
	return o.Username == nil && o.Locale == nil && o.Bio == nil
}
