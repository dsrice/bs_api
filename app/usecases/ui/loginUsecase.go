package ui

import (
	"app/controllers/requestparameter"
	"app/entities"
)

type LoginUsecase interface {
	Validate(param requestparameter.Login) error
	GetUser(loginID string) (*entities.UserEntity, error)
}