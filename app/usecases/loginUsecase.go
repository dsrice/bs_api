package usecases

import (
	"app/controllers/rqp"
	"app/entities"
	"app/infra/logger"
	"app/repositories/ri"
	"app/usecases/ui"
	"fmt"
)

type loginUsecaseImp struct {
	ur ri.UserRepository
}

func NewLoginUsecase(repo ri.InRepository) ui.LoginUsecase {
	return &loginUsecaseImp{
		ur: repo.UserRepo,
	}
}

func (u *loginUsecaseImp) Validate(param rqp.Login) error {

	if len(param.LoginID) == 0 || len(param.Password) == 0 {
		err := fmt.Errorf("ログインIDもしくはパスワードがありません")
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (u *loginUsecaseImp) GetUser(loginID string) (*entities.UserEntity, error) {
	user, err := u.ur.GetUser(loginID)

	if err != nil {
		return nil, err
	}

	return user, nil
}