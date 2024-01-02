package rmock

import (
	"app/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (_m *UserRepositoryMock) GetUser(loginID string) (*entities.UserEntity, error) {
	ret := _m.Called(loginID)
	return ret.Get(0).(*entities.UserEntity), ret.Error(1)
}