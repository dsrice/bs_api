package ci

import "github.com/labstack/echo/v4"

type LoginController interface {
	Get(c echo.Context) error
}