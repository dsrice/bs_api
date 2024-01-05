package usecases

import (
	"app/controllers/rqp"
	"app/entities"
	"app/infra/logger"
	ri "app/repositories/ri"
	"app/usecases/ui"
	"fmt"
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

	logger.Debug("RegistValidate Usecase Start")
	return nil
}

func (uc *userUsecaseImp) CheckUser(loginID string) (*entities.UserEntity, error) {
	um, err := uc.uRepo.GetUser(loginID)

	if err != nil {
		return nil, err
	}

	return um, nil
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