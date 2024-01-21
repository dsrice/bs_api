package rmock

import (
	"app/entities/token"
	"app/entities/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type TokenRepositoryMock struct {
	mock.Mock
}

func (m *TokenRepositoryMock) SetToken(user user.Entity, c echo.Context) (*token.Entity, error) {
	ret := m.Called(user, c)
	return ret.Get(0).(*token.Entity), ret.Error(1)
}

func (m *TokenRepositoryMock) SearchUser(token string, c echo.Context) (*user.Entity, error) {
	ret := m.Called(token, c)
	return ret.Get(0).(*user.Entity), ret.Error(1)
}