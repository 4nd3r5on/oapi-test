// Package users provides application layer for users
package users

type Service struct {
	db DB
}

func NewService(db DB) *Service {
	return &Service{db: db}
}
