package ui

import (
	"app/controllers/rqp"
	"app/entities"
)

type UserUsecase interface {
	RegistValidate(param rqp.RegistUser) error
	RegistUser(user *entities.UserEntity) error
}