package repositories_test

import (
	"app/entities/user"
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
	user := user.Entity{UserID: 1}
	t, err := s.repo.SetToken(user)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), t.User.UserID, 1)
}

func TestSetTokenSuite(t *testing.T) {
	suite.Run(t, new(SetTokenSuite))
}

type SearchUserSuite struct {
	suite.Suite
	repo ri.TokenRepository
}

func (s *SearchUserSuite) SetupSuite() {
	conn := connection.NewConnection()
	s.repo = repositories.NewTokenRepository(conn)
}

func (s *SearchUserSuite) TestSuccess() {
	u := user.Entity{UserID: 1}
	t, err := s.repo.SetToken(u)

	result, err := s.repo.SearchUser(t.Token)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result.LoginID, "test")
}

func TestSearchUserSuite(t *testing.T) {
	suite.Run(t, new(SearchUserSuite))
}