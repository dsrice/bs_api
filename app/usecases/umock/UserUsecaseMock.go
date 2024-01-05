package umock

import (
	"app/entities"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *LoginUsecaseMock) RegistUser(user *entities.UserEntity) error {
	ret := m.Called(user)
	return ret.Error(0)
}

func (m *LoginUsecaseMock) CheckUser(loginID string) (*entities.UserEntity, error) {
	ret := m.Called(loginID)
	return ret.Get(0).(*entities.UserEntity), ret.Error(1)
}