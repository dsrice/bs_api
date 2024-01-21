package ri

import (
	"app/entities/user"
)

type UserRepository interface {
	GetUser(us *user.Search) ([]*user.Entity, error)
	RegistUser(user user.Entity) error
}