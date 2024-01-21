package ui

import (
	"app/controllers/rqp"
	"app/entities/user"
	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	RegistValidate(param rqp.RegistUser) error
	CheckUser(loginID string, c echo.Context) (*user.Entity, error)
	RegistUser(user *user.Entity, c echo.Context) error
	GetUsers(uc *user.Search, c echo.Context) ([]*user.Entity, error)
}