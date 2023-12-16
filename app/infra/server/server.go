package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	s := Server{}

	return &s
}

func (s *Server) Start() {
	s.echo = echo.New()
	s.routing()
	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}

func (s *Server) routing() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hellow World")
	})
}
