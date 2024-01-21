package umock

import (
	"app/controllers/rqp"
	"app/entities/user"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) RegistValidate(param rqp.RegistUser) error {
	ret := m.Called(param)
	return ret.Error(0)
}

func (m *UserUsecaseMock) CheckUser(loginID string) (*user.Entity, error) {
	ret := m.Called(loginID)
	return ret.Get(0).(*user.Entity), ret.Error(1)
}

func (m *UserUsecaseMock) RegistUser(user *user.Entity) error {
	ret := m.Called(user)
	return ret.Error(0)
}

func (m *UserUsecaseMock) GetUsers(uc *user.Search) ([]*user.Entity, error) {
	ret := m.Called(uc)
	return ret.Get(0).([]*user.Entity), ret.Error(1)
}