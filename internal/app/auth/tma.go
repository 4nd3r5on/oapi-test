package auth

import (
	"context"
	"time"

	"github.com/4nd3rs0n/oapi-test/pkg/errs"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// TMA stands for Telegram Mini App auth
// Users telegram init data to verify user
type TMA struct {
	BotToken string
}

func (auth *TMA) Verify(_ context.Context, _, initData string, _ []string) (*ClientData, error) {
	parsedData, err := initdata.Parse(initData)
	if err != nil {
		err = errs.Mark(errs.ErrInvalidArgument)
		err = errs.SafeMsg(err, "failed to parse TMA authorization token")
		return nil, err
	}
	err = initdata.Validate(initData, auth.BotToken, 24*time.Hour)
	if err != nil {
		err = errs.Mark(errs.ErrPermissionDenied)
		err = errs.SafeMsg(err, "TMA initdata validation failed")
		return nil, err
	}
	return &ClientData{
		Method:     MethodTMA,
		TgInitData: &parsedData,
	}, nil
}
