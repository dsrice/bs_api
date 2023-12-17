package controllers

import (
	"app/controllers/ci"
	"app/repositories/ri"
	"github.com/labstack/echo/v4"
	"net/http"
)

type loginControllerImp struct {
	repo ri.UserRepository
}

func NewLoginController(repo ri.UserRepository) ci.LoginController {
	return &loginControllerImp{
		repo: repo,
	}
}

func (ct *loginControllerImp) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hellow World")
}