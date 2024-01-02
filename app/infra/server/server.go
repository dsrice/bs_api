package server

import (
	"app/controllers/ci"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo  *echo.Echo
	Login ci.LoginController
}

func NewServer(s ci.InController) *Server {
	return &Server{
		Login: s.Login,
	}
}

func (s *Server) Start() {
	s.echo = echo.New()
	s.routing()

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}