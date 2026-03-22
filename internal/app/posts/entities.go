package posts

import "time"

type Post struct {
	ID        int64
	Text      string
	CreatedAt time.Time
	EditedAt  *time.Time
}
