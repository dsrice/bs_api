package ui

import "app/entities"

type LoginUsecase interface {
	GetUser(loginID string) (*entities.UserEntity, error)
}