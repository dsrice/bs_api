package controllers

import (
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/infra/errormessage"
	"app/infra/logger"
	"app/infra/response"
	"app/usecases/ui"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
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
		if err.Error() == strconv.Itoa(errormessage.UsedLoginID) {
			return response.ErrorResponse(c, errormessage.UsedLoginID)
		} else {
			return response.ErrorResponse(c, errormessage.BadRequest)
		}
	}

	user := param.ConvertEntity()

	err = ct.userUC.RegistUser(&user)

	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.SystemError)
	}

	logger.Debug("RegistUser API End")
	return nil
}