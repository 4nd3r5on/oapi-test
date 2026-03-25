package users

import (
	"time"

	"github.com/google/uuid"
)

// Domain entities
type (
	// User domain entity
	User struct {
		ID        uuid.UUID
		Username  string
		Locale    string
		Bio       string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	// UserIdentityTelegram used for Telegram Mini App auth
	// over the tma init data
	UserIdentityTelegram struct {
		UserID     uuid.UUID // PK
		TelegramID int64
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	UserIdentityOIDC struct {
		UserID    uuid.UUID // PK
		Subject   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

// Options
type (
	// NewUserOpts is the data that should be provided
	// in a request to create a new user
	NewUserOpts struct {
		ID       uuid.UUID
		Username string
		Locale   string
		Bio      string
	}
	// UpdateUserOpts is the data that can be updated on a user's own profile.
	// All fields are optional; at least one must be set.
	UpdateUserOpts struct {
		Username *string
		Locale   *string
		Bio      *string
	}

	NewUserIdentityTelegramOpts struct {
		UserID     uuid.UUID
		TelegramID int64
	}
	NewUserIdentityOIDCOpts struct {
		UserID  uuid.UUID
		Subject string
	}
)

// Result variants
type (
	// PublicUser contains public user info which
	// any other user can see (by id or username)
	PublicUser struct {
		ID       uuid.UUID
		Username string
		Bio      string
	}
	// PrivateUser contains private user info which
	// only the user himself can see
	PrivateUser struct {
		PublicUser
		// can be extended with private data in  the future
		Locale    string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
