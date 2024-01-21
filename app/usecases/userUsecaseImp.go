package usecases

import (
	"app/controllers/rqp"
	"app/entities/user"
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

func (uc *userUsecaseImp) CheckUser(loginID string) (*user.Entity, error) {
	us := user.Search{
		LoginID: &loginID,
	}
	um, err := uc.uRepo.GetUser(&us)

	if err != nil {
		return nil, err
	}

	if len(um) != 1 {
		err = fmt.Errorf(fmt.Sprintf("ユーザーが見つかりませんでいた。 該当数:%d", len(um)))
		logger.Error(err.Error())
		return nil, err
	}

	return um[0], nil
}

func (uc *userUsecaseImp) RegistUser(user *user.Entity) error {
	logger.Debug("RegistUser Usecase Start")
	err := uc.uRepo.RegistUser(*user)

	if err != nil {
		return err
	}

	logger.Debug("RegistUser Usecase End")
	return nil
}

func (uc *userUsecaseImp) GetUsers(us *user.Search) ([]*user.Entity, error) {
	logger.Debug("GetUsers Usecase Start")
	ul, err := uc.uRepo.GetUser(us)

	if err != nil {
		return nil, err
	}

	logger.Debug("GetUsers Usecase End")
	return ul, nil
}