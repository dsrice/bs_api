package umock

import (
	"app/controllers/rqp"
	"app/entities/token"
	"app/entities/user"
	"github.com/stretchr/testify/mock"
)

type LoginUsecaseMock struct {
	mock.Mock
}

func (m *LoginUsecaseMock) Validate(param rqp.Login) error {
	ret := m.Called(param)
	return ret.Error(0)
}

func (m *LoginUsecaseMock) GetUser(loginID string) (*user.Entity, error) {
	ret := m.Called(loginID)
	return ret.Get(0).(*user.Entity), ret.Error(1)
}

func (m *LoginUsecaseMock) GetToken(user *user.Entity) (*token.Entity, error) {
	ret := m.Called(user)
	return ret.Get(0).(*token.Entity), ret.Error(1)
}