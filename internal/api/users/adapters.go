package users

import (
	appusers "github.com/4nd3rs0n/oapi-test/internal/app/users"
	"github.com/4nd3rs0n/oapi-test/pkg/api"
)

func privateUserToAPI(u *appusers.PrivateUser) *api.PrivUser {
	return &api.PrivUser{
		ID:        api.UUID(u.ID),
		Username:  u.Username,
		Bio:       u.Bio,
		Locale:    u.Locale,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func publicUserToAPI(u *appusers.PublicUser) *api.PubUser {
	return &api.PubUser{
		ID:       api.UUID(u.ID),
		Username: u.Username,
		Bio:      u.Bio,
	}
}
