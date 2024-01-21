package ri

import (
	"app/entities/token"
	"app/entities/user"
	"github.com/labstack/echo/v4"
)

type TokenRepository interface {
	SetToken(user user.Entity, c echo.Context) (*token.Entity, error)
	SearchUser(token string, c echo.Context) (*user.Entity, error)
}