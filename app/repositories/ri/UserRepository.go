package ri

import "app/entities"

type UserRepository interface {
	GetUser(loginID, name, mail *string) ([]*entities.UserEntity, error)
	RegistUser(user entities.UserEntity) error
}