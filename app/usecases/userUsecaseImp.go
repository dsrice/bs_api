package usecases

import (
	"app/entities"
	"app/infra/logger"
	ri "app/repositories/ri"
	"app/usecases/ui"
)

type userUsecaseImp struct {
	uRepo ri.UserRepository
}

func NewUserUsecase(repo ri.InRepository) ui.UserUsecase {
	return &userUsecaseImp{
		uRepo: repo.UserRepo,
	}
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