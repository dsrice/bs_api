package controllers

import (
	"app/controllers/ci"
	"app/controllers/requestparameter"
	"app/infra/logger"
	"app/infra/response"
	"app/usecases/ui"
	"github.com/labstack/echo/v4"
	"net/http"
)

type loginControllerImp struct {
	login ui.LoginUsecase
}

func NewLoginController(uc ui.InUsecase) ci.LoginController {
	return &loginControllerImp{
		login: uc.Login,
	}
}

func (ct *loginControllerImp) Login(c echo.Context) error {
	param := requestparameter.Login{}
	if err := c.Bind(param); err != nil {
		return err
	}

	_, err := ct.login.GetUser(param.LoginID)

	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, "NoUser")
	}

	return c.String(http.StatusOK, "Hellow World")
}