package main

import (
	"app/di"
	"app/infra/server"
	"go.uber.org/dig"
)

func main() {
	c := dig.New()

	di.BuildContainer(c)
	err := c.Invoke(func(s *server.Server) {
		s.Start()
	})

	if err != nil {
		panic(err)
	}
}