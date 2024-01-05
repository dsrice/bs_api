package server

import (
	"app/controllers/ci"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	echo  *echo.Echo
	Login ci.LoginController
	User  ci.UserController
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewServer(s ci.InController) *Server {
	return &Server{
		Login: s.Login,
		User:  s.UserController,
	}
}

func (s *Server) Start() {
	s.echo = echo.New()
	s.echo.Validator = &CustomValidator{Validator: validator.New()}
	s.routing()

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}