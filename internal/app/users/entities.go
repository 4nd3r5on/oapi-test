package users

import (
	"time"

	"github.com/google/uuid"
)

// User domain entity
type User struct {
	ID        uuid.UUID
	Username  string
	Locale    string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUserOpts is the data that should be provided
// in a request to create a new user
type NewUserOpts struct {
	ID       uuid.UUID
	Username string
	Locale   string
	Bio      string
}

// PubUser contains public user info which
// any other user can see (by id or username)
type PubUser struct {
	ID       uuid.UUID
	Username string
	Bio      string
}

// PrivUser contains private user info which
// only the user himself can see
type PrivUser struct {
	PubUser
	// can be extended with private data in  the future
	Locale    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UpdateUserOpts is the data that can be updated on a user's own profile.
// All fields are optional; at least one must be set.
type UpdateUserOpts struct {
	Username *string
	Locale   *string
	Bio      *string
}
