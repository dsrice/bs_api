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
	tr ri.TokenRepository
}

func NewLoginUsecase(repo ri.InRepository) ui.LoginUsecase {
	return &loginUsecaseImp{
		ur: repo.UserRepo,
		tr: repo.TokenRepository,
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

func (u *loginUsecaseImp) GetToken(user *entities.UserEntity) (*entities.TokenEntity, error) {
	token, err := u.tr.SetToken(*user)

	if err != nil {
		return nil, err
	}

	return token, nil
}