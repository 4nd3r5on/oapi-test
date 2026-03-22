package messages

import "context"

type App interface {
	NewMessage(ctx context.Context) error
	LoadMessages(ctx context.Context) error
	LoadUserMessages(ctx context.Context) error
	DeleteMessage(ctx context.Context) error
	EditMessage(ctx context.Context) error
}
