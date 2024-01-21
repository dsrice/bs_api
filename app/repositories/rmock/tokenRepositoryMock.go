package rmock

import (
	"app/entities/token"
	"app/entities/user"
	"github.com/stretchr/testify/mock"
)

type TokenRepositoryMock struct {
	mock.Mock
}

func (m *TokenRepositoryMock) SetToken(user user.Entity) (*token.Entity, error) {
	ret := m.Called(user)
	return ret.Get(0).(*token.Entity), ret.Error(1)
}

func (m *TokenRepositoryMock) SearchUser(token string) (*user.Entity, error) {
	ret := m.Called(token)
	return ret.Get(0).(*user.Entity), ret.Error(1)
}