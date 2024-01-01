package controllers

import (
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/controllers/rsp"
	"app/infra/errormessage"
	"app/infra/logger"
	"app/infra/response"
	"app/usecases/ui"
	"github.com/labstack/echo/v4"
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
	logger.Debug("login API Start")
	param := rqp.Login{}
	if err := c.Bind(&param); err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.BadRequest)
	}

	err := ct.login.Validate(param)

	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.BadRequest)
	}

	user, err := ct.login.GetUser(param.LoginID)

	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.FailAuth)
	}

	token, err := ct.login.GetToken(user)
	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.SystemError)
	}

	res := rsp.Login{}
	res.ConvertResponse(token)

	logger.Debug("login API End")
	return response.SccessResponse(c, res)
}