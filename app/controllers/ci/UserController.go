package ci

import "github.com/labstack/echo/v4"

type UserController interface {
	RegistUser(c echo.Context) error
	GetUser(c echo.Context) error
}