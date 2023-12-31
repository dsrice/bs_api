package ci

import "github.com/labstack/echo/v4"

type LoginController interface {
	Login(c echo.Context) error
}