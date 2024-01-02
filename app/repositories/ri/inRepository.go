package ri

import "go.uber.org/dig"

type InRepository struct {
	dig.In
	UserRepo        UserRepository
	TokenRepository TokenRepository
}