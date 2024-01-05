package ri

import "app/entities"

type UserRepository interface {
	GetUser(loginID string) (*entities.UserEntity, error)
	RegistUser(user entities.UserEntity) error
}