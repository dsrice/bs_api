package controllers

import (
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/controllers/rsp"
	"app/entities/user"
	"app/infra/errormessage"
	"app/infra/logger"
	"app/infra/response"
	"app/infra/server"
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

	ue := param.ConvertUserEntity()

	err = ct.userUC.RegistUser(&ue)

	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.SystemError)
	}

	res := rsp.User{}
	res.ConvertResponse(&ue)

	logger.Debug("RegistUser API End")
	return response.SccessResponse(c, res)
}

func (ct *userControllerImp) GetUser(c echo.Context) error {
	logger.Debug("GetUser API Start")

	sess := server.GetSession(c)
	t := sess.Values["user"]
	fmt.Println(t)

	uc := user.Search{}
	ul, err := ct.userUC.GetUsers(&uc)

	if err != nil {
		logger.Error(err.Error())
		return response.ErrorResponse(c, errormessage.SystemError)
	}

	var rl []*rsp.User
	for _, u := range ul {
		r := rsp.User{}
		r.ConvertResponse(u)
		rl = append(rl, &r)
	}

	res := rsp.GetUser{Users: rl}

	logger.Debug("GetUser API End")
	return response.SccessResponse(c, res)
}