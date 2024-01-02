package di

import (
	"app/infra/database/connection"
	"app/infra/logger"
	"app/infra/server"
	"go.uber.org/dig"
)

func BuildContainer(c *dig.Container) {
	setProvide(c, server.NewServer)
	setProvide(c, connection.NewConnection)
	provideController(c)
	provideUsecase(c)
	provideRepository(c)
}

func setProvide(c *dig.Container, i interface{}) {
	if err := c.Provide(i); err != nil {
		logger.Error(err.Error())
	}
}