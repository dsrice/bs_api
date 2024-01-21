package ri

import (
	"app/entities/user"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	GetUser(us *user.Search, c echo.Context) ([]*user.Entity, error)
	RegistUser(user user.Entity, c echo.Context) error
}