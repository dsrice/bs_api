package repositories_test

import (
	"app/entities/user"
	"app/infra/database/connection"
	"app/repositories"
	"app/repositories/ri"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
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

func (s *GetUserSuite) TestSuccessALL() {
	us := user.Search{}
	ul, err := s.repo.GetUser(&us)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), ul[0].LoginID, "test")
	assert.Equal(s.T(), ul[0].Name, "test")
}

func (s *GetUserSuite) TestSuccessLoginID() {
	loginID := "test"
	us := user.Search{
		LoginID: &loginID,
	}
	ul, err := s.repo.GetUser(&us)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), ul[0].LoginID, "test")
	assert.Equal(s.T(), ul[0].Name, "test")
}

func (s *GetUserSuite) TestSuccessName() {
	name := "test"
	us := user.Search{
		Name: &name,
	}
	ul, err := s.repo.GetUser(&us)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), ul[0].LoginID, "test")
	assert.Equal(s.T(), ul[0].Name, "test")
}

func (s *GetUserSuite) TestSuccessMail() {
	mail := "test@test"
	us := user.Search{
		Mail: &mail,
	}

	ul, err := s.repo.GetUser(&us)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), ul[0].LoginID, "test")
	assert.Equal(s.T(), ul[0].Name, "test")
}

func TestGetUserSuite(t *testing.T) {
	suite.Run(t, new(GetUserSuite))
}

type RegistUserSuite struct {
	suite.Suite
	repo ri.UserRepository
}

func (s *RegistUserSuite) SetupSuite() {
	conn := connection.NewConnection()
	s.repo = repositories.NewUserRepository(conn)

	user := user.Entity{
		UserID:   11,
		LoginID:  "testtest",
		Name:     "test",
		Password: "test",
		Mail:     "test@test",
	}

	m := user.ConvertUserModel()
	_, err := m.Delete(context.Background(), conn.Conn)
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *RegistUserSuite) Test_01_Success() {
	user := user.Entity{
		UserID:   11,
		LoginID:  "testtest",
		Name:     "test",
		Password: "test",
		Mail:     "test@test",
	}

	err := s.repo.RegistUser(user)

	assert.Nil(s.T(), err)
}

func (s *RegistUserSuite) Test_02_Failed() {
	user := user.Entity{
		UserID:   11,
		LoginID:  "testtest",
		Name:     "test",
		Password: "test",
		Mail:     "test@test",
	}

	err := s.repo.RegistUser(user)

	assert.NotNil(s.T(), err)
}

func TestRegistUserSuite(t *testing.T) {
	suite.Run(t, new(RegistUserSuite))
}