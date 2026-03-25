package sessions

import (
	"context"
)

type Cache interface {
	StoreSession(ctx context.Context, token string, session Session) error
	GetSession(ctx context.Context, token string) (*Session, bool, error)
	DeleteSession(ctx context.Context, token string) error
}

