package di

import (
	"app/repositories"
	"go.uber.org/dig"
)

func provideRepository(c *dig.Container) {
	setProvide(c, repositories.NewUserRepository)
}