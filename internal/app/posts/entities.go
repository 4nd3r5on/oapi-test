package posts

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        int64
	AuthorID  uuid.UUID
	Title     string
	Text      string
	CreatedAt time.Time
	EditedAt  *time.Time
}
