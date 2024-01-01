package repositories_test

import (
	"app/entities"
	"app/infra/database/connection"
	"app/repositories"
	"app/repositories/ri"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SetTokenSuite struct {
	suite.Suite
	repo ri.TokenRepository
}

func (s *SetTokenSuite) SetupSuite() {
	conn := connection.NewConnection()
	s.repo = repositories.NewTokenRepository(conn)
}

func (s *SetTokenSuite) TestSuccess() {
	user := entities.UserEntity{UserID: 1}
	t, err := s.repo.SetToken(user)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), t.User.UserID, 1)
}

func TestSetTokenSuite(t *testing.T) {
	suite.Run(t, new(SetTokenSuite))
}