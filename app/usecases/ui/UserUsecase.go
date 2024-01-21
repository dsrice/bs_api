package ui

import (
	"app/controllers/rqp"
	"app/entities/user"
)

type UserUsecase interface {
	RegistValidate(param rqp.RegistUser) error
	CheckUser(loginID string) (*user.Entity, error)
	RegistUser(user *user.Entity) error
	GetUsers(uc *user.Search) ([]*user.Entity, error)
}