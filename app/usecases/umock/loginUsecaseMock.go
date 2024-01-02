package umock

import (
	"app/controllers/rqp"
	"app/entities"
	"github.com/stretchr/testify/mock"
)

type LoginUsecaseMock struct {
	mock.Mock
}

func (m *LoginUsecaseMock) Validate(param rqp.Login) error {
	ret := m.Called(param)
	return ret.Error(0)
}

func (m *LoginUsecaseMock) GetUser(loginID string) (*entities.UserEntity, error) {
	ret := m.Called(loginID)
	return ret.Get(0).(*entities.UserEntity), ret.Error(1)
}

func (m *LoginUsecaseMock) GetToken(user *entities.UserEntity) (*entities.TokenEntity, error) {
	ret := m.Called(user)
	return ret.Get(0).(*entities.TokenEntity), ret.Error(1)
}