package rmock

import (
	"app/entities"
	"github.com/stretchr/testify/mock"
)

type TokenRepositoryMock struct {
	mock.Mock
}

func (_m *TokenRepositoryMock) SetToken(user entities.UserEntity) (*entities.TokenEntity, error) {
	ret := _m.Called(user)
	return ret.Get(0).(*entities.TokenEntity), ret.Error(1)
}