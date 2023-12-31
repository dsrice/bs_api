package ui

import (
	"app/controllers/rqp"
	"app/entities"
)

type LoginUsecase interface {
	Validate(param rqp.Login) error
	GetUser(loginID string) (*entities.UserEntity, error)
}