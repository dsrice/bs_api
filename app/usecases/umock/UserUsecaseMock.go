package umock

import (
	"app/controllers/rqp"
	"app/entities"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) RegistValidate(param rqp.RegistUser) error {
	ret := m.Called(param)
	return ret.Error(0)
}

func (m *UserUsecaseMock) CheckUser(loginID string) (*entities.UserEntity, error) {
	ret := m.Called(loginID)
	return ret.Get(0).(*entities.UserEntity), ret.Error(1)
}

func (m *UserUsecaseMock) RegistUser(user *entities.UserEntity) error {
	ret := m.Called(user)
	return ret.Error(0)
}