package ri

import (
	"app/entities/token"
	"app/entities/user"
)

type TokenRepository interface {
	SetToken(user user.Entity) (*token.Entity, error)
	SearchUser(token string) (*user.Entity, error)
}