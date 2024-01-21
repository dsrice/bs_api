package umock

import (
	"app/controllers/rqp"
	"app/entities/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) RegistValidate(param rqp.RegistUser) error {
	ret := m.Called(param)
	return ret.Error(0)
}

func (m *UserUsecaseMock) CheckUser(loginID string, c echo.Context) (*user.Entity, error) {
	ret := m.Called(loginID, c)
	return ret.Get(0).(*user.Entity), ret.Error(1)
}

func (m *UserUsecaseMock) RegistUser(user *user.Entity, c echo.Context) error {
	ret := m.Called(user, c)
	return ret.Error(0)
}

func (m *UserUsecaseMock) GetUsers(uc *user.Search, c echo.Context) ([]*user.Entity, error) {
	ret := m.Called(uc, c)
	return ret.Get(0).([]*user.Entity), ret.Error(1)
}