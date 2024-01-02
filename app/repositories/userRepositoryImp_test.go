package repositories_test

import (
	"app/infra/database/connection"
	"app/repositories"
	"app/repositories/ri"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GetUserSuite struct {
	suite.Suite
	repo ri.UserRepository
}

func (s *GetUserSuite) SetupSuite() {
	conn := connection.NewConnection()
	s.repo = repositories.NewUserRepository(conn)
}

func (s *GetUserSuite) TestSuccess() {
	u, err := s.repo.GetUser("test")

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), u.LoginID, "test")
}

func (s *GetUserSuite) TestFailNoUser() {
	u, err := s.repo.GetUser("")

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "対象ユーザが見つかりませんでした")
	assert.Nil(s.T(), u)
}

func TestGetUserSuite(t *testing.T) {
	suite.Run(t, new(GetUserSuite))
}