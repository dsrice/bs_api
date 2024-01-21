package rmock

import (
	"app/entities/token"
	"app/entities/user"
	"github.com/stretchr/testify/mock"
)

type TokenRepositoryMock struct {
	mock.Mock
}

func (_m *TokenRepositoryMock) SetToken(user user.Entity) (*token.Entity, error) {
	ret := _m.Called(user)
	return ret.Get(0).(*token.Entity), ret.Error(1)
}