package ri

import "app/entities"

type TokenRepository interface {
	SetToken(user entities.UserEntity) (*entities.TokenEntity, error)
}