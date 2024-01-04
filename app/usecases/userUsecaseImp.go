package usecases

import (
	"app/controllers/rqp"
	"app/entities"
	"app/infra/errormessage"
	"app/infra/logger"
	ri "app/repositories/ri"
	"app/usecases/ui"
	"fmt"
	"strconv"
)

type userUsecaseImp struct {
	uRepo ri.UserRepository
}

func NewUserUsecase(repo ri.InRepository) ui.UserUsecase {
	return &userUsecaseImp{
		uRepo: repo.UserRepo,
	}
}

func (uc *userUsecaseImp) RegistValidate(param rqp.RegistUser) error {
	logger.Debug("RegistValidate Usecase Start")

	if len(param.LoginID) == 0 || len(param.Password) == 0 || len(param.Name) == 0 {
		return fmt.Errorf("必須項目が未設定")
	}

	user, err := uc.uRepo.GetUser(param.LoginID)

	if err != nil {
		return err
	}

	if user != nil {
		err = fmt.Errorf(strconv.Itoa(errormessage.UsedLoginID))
		logger.Error(err.Error())
		return err
	}

	logger.Debug("RegistValidate Usecase Start")
	return nil
}

func (uc *userUsecaseImp) RegistUser(user *entities.UserEntity) error {
	logger.Debug("RegistUser Usecase Start")
	err := uc.uRepo.RegistUser(*user)

	if err != nil {
		return err
	}

	logger.Debug("RegistUser Usecase End")
	return nil
}