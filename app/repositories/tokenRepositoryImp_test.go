package repositories_test

import (
	"app/entities/user"
	"app/infra/database/connection"
	"app/repositories"
	"app/repositories/ri"
	"app/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type SetTokenSuite struct {
	suite.Suite
	repo   ri.TokenRepository
	tester tester.Tester
}

func (s *SetTokenSuite) SetupSuite() {
	conn := connection.NewConnection()
	s.repo = repositories.NewTokenRepository(conn)
	s.tester = tester.CreateContext(http.MethodPost, "/", nil)
}

func (s *SetTokenSuite) TestSuccess() {
	u := user.Entity{UserID: 1}
	t, err := s.repo.SetToken(u, s.tester.Context)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), t.User.UserID, 1)
}

func TestSetTokenSuite(t *testing.T) {
	suite.Run(t, new(SetTokenSuite))
}

type SearchUserSuite struct {
	suite.Suite
	repo   ri.TokenRepository
	tester tester.Tester
}

func (s *SearchUserSuite) SetupSuite() {
	conn := connection.NewConnection()
	s.repo = repositories.NewTokenRepository(conn)
	s.tester = tester.CreateContext(http.MethodPost, "/", nil)
}

func (s *SearchUserSuite) TestSuccess() {
	u := user.Entity{UserID: 1}
	t, err := s.repo.SetToken(u, s.tester.Context)

	result, err := s.repo.SearchUser(t.Token, s.tester.Context)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result.LoginID, "test")
}

func TestSearchUserSuite(t *testing.T) {
	suite.Run(t, new(SearchUserSuite))
}