package server

import (
	"app/controllers/ci"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Server struct {
	echo  *echo.Echo
	Login ci.LoginController
}

type inServer struct {
	dig.In
	Login ci.LoginController
}

func NewServer(s inServer) *Server {
	return &Server{
		Login: s.Login,
	}
}

func (s *Server) Start() {
	s.echo = echo.New()
	s.routing()

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}

func (s *Server) routing() {
	s.echo.GET("/", s.Login.Get)
}