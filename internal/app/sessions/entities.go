package sessions

import "github.com/google/uuid"

const KeyLen = 32

type Session struct {
	UserID uuid.UUID
}
