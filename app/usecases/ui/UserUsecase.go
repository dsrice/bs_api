package ui

import "app/entities"

type UserUsecase interface {
	RegistUser(user *entities.UserEntity) error
}