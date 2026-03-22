// Package auth provides auth methods/verifires
package auth

import (
	"context"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type Method int

const (
	MethodUnknown Method = iota

	MethodTMA     // Telegram Mini App
	MethodTMB     // Trust Me Bro
	MethodSession // Internal aplication sessions
)

var methodToString = map[Method]string{
	MethodUnknown: "unknown",
	MethodTMA:     "tma",
	MethodTMB:     "tmb",
	MethodSession: "session",
}

func (m Method) String() string {
	if s, ok := methodToString[m]; ok {
		return s
	}
	return "unknown"
}

type ClientData struct {
	Method     Method
	TgInitData *initdata.InitData
	TMBData    *TMBData
}

type VerifierFunc func(ctx context.Context, data string) (*ClientData, error)
