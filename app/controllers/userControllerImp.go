package controllers

import (
	"app/controllers/ci"
	"app/infra/logger"
	"app/usecases/ui"
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

	logger.Debug("RegistUser API End")
	return nil
}