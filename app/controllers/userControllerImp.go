package controllers

import (
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/controllers/rsp"
	"app/infra/errormessage"
	"app/infra/logger"
	"app/infra/response"
	"app/usecases/ui"
	"fmt"
	"github.com/labstack/echo/v4"
)

type userControllerImp struct {
	userUC ui.UserUsecase
}

func NewUserController(uc ui.InUsecase) ci.UserController {
	return &userControllerImp{
		userUC: uc.UserUsecase,
	}
}

func (ct *userControllerImp) RegistUser(c echo.Context) error {
	logger.Debug("RegistUser API Start")
	param := rqp.RegistUser{}
	if err := c.Bind(&param); err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.BadRequest)
	}

	if ev := c.Validate(param); ev != nil {
		logger.Error(fmt.Sprintf("param: %s", param))
		logger.Error(ev.Error())
		return response.ErrorResponse(c, errormessage.BadRequest)
	}

	err := ct.userUC.RegistValidate(param)
	if err != nil {
		return response.ErrorResponse(c, errormessage.BadRequest)
	}

	uc, err := ct.userUC.CheckUser(param.LoginID)

	if err != nil {
		return response.ErrorResponse(c, errormessage.SystemError)
	} else if uc != nil {
		return response.ErrorResponse(c, errormessage.UsedLoginID)
	}

	user := param.ConvertEntity()

	err = ct.userUC.RegistUser(&user)

	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.SystemError)
	}

	res := rsp.User{}
	res.ConvertResponse(&user)

	logger.Debug("RegistUser API End")
	return response.SccessResponse(c, res)
}