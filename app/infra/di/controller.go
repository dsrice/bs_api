package di

import (
	"app/controllers"
	"go.uber.org/dig"
)

func provideController(c *dig.Container) *dig.Container {
	setProvide(c, controllers.NewLoginController)

	return c
}