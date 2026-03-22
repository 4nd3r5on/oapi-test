package auth

import (
	"context"
	"time"

	"github.com/4nd3r5on/errs"
	"github.com/google/uuid"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type AuthTMARepo interface {
	GetUserIDByTgID(ctx context.Context, tgID int64) (uuid.UUID, error)
}

type TMA struct {
	BotToken string
	Repo     AuthTMARepo
}

func (auth *TMA) Verify(ctx context.Context, data string) (*ClientData, error) {
	parsedData, err := initdata.Parse(data)
	if err != nil {
		return nil, errs.F().
			Message("failed to parse TMA authorization token").
			Mark(errs.ErrInvalidArgument).Err()
	}
	err = initdata.Validate(data, auth.BotToken, 24*time.Hour)
	if err != nil {
		return nil, errs.F().
			Message("TMA initdata validation failed").
			Mark(errs.ErrPermissionDenied).Err()
	}
	return &ClientData{
		Method:     MethodTMA,
		TgInitData: &parsedData,
	}, nil
}
