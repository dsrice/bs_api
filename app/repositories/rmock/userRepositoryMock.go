package rmock

import (
	"app/entities/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (_m *UserRepositoryMock) GetUser(us *user.Search, c echo.Context) ([]*user.Entity, error) {
	ret := _m.Called(us, c)
	return ret.Get(0).([]*user.Entity), ret.Error(1)
}

func (_m *UserRepositoryMock) RegistUser(user user.Entity, c echo.Context) error {
	ret := _m.Called(user, c)
	return ret.Error(0)
}