package usecases

import (
	"app/entities"
	"app/repositories/ri"
	"app/usecases/ui"
)

type loginUsecaseImp struct {
	ur ri.UserRepository
}

func NewLoginUsecase(repo ri.InRepository) ui.LoginUsecase {
	return &loginUsecaseImp{
		ur: repo.UserRepo,
	}
}

func (u *loginUsecaseImp) GetUser(loginID string) (*entities.UserEntity, error) {
	user, err := u.ur.GetUser(loginID)

	if err != nil {
		return nil, err
	}

	return user, nil
}