package di

import (
	"app/usecases"
	"go.uber.org/dig"
)

func provideUsecase(c *dig.Container) {
	setProvide(c, usecases.NewLoginUsecase)
	setProvide(c, usecases.NewUserUsecase)
}