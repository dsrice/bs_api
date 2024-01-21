package rmock

import (
	"app/entities/user"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (_m *UserRepositoryMock) GetUser(us *user.Search) ([]*user.Entity, error) {
	ret := _m.Called(us)
	return ret.Get(0).([]*user.Entity), ret.Error(1)
}

func (_m *UserRepositoryMock) RegistUser(user user.Entity) error {
	ret := _m.Called(user)
	return ret.Error(0)
}