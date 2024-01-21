package ui

import (
	"app/controllers/rqp"
	"app/entities/token"
	"app/entities/user"
)

type LoginUsecase interface {
	Validate(param rqp.Login) error
	GetUser(loginID string) (*user.Entity, error)
	GetToken(user *user.Entity) (*token.Entity, error)
}