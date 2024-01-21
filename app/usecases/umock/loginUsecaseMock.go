package umock

import (
	"app/controllers/rqp"
	"app/entities/token"
	"app/entities/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type LoginUsecaseMock struct {
	mock.Mock
}

func (m *LoginUsecaseMock) Validate(param rqp.Login) error {
	ret := m.Called(param)
	return ret.Error(0)
}

func (m *LoginUsecaseMock) GetUser(loginID string, c echo.Context) (*user.Entity, error) {
	ret := m.Called(loginID, c)
	return ret.Get(0).(*user.Entity), ret.Error(1)
}

func (m *LoginUsecaseMock) GetToken(user *user.Entity, c echo.Context) (*token.Entity, error) {
	ret := m.Called(user, c)
	return ret.Get(0).(*token.Entity), ret.Error(1)
}