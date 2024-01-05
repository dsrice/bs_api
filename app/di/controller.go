package di

import (
	"app/controllers"
	"go.uber.org/dig"
)

func provideController(c *dig.Container) {
	setProvide(c, controllers.NewLoginController)
	setProvide(c, controllers.NewUserController)
}