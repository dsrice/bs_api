package ui

import (
	"app/controllers/rqp"
	"app/entities/token"
	"app/entities/user"
	"github.com/labstack/echo/v4"
)

type LoginUsecase interface {
	Validate(param rqp.Login) error
	GetUser(loginID string, c echo.Context) (*user.Entity, error)
	GetToken(user *user.Entity, c echo.Context) (*token.Entity, error)
}