package repositories_test

import (
	"app/entities"
	"app/infra/database/connection"
	"app/repositories"
	"app/repositories/ri"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/boil"
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
	boil.DebugMode = true
}

func (s *GetUserSuite) TestSuccess() {
	u, err := s.repo.GetUser("test")

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), u.LoginID, "test")
}

func (s *GetUserSuite) TestFailNoUser() {
	u, err := s.repo.GetUser("")

	assert.Nil(s.T(), err)
	assert.Nil(s.T(), u)
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

	user := entities.UserEntity{
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
	user := entities.UserEntity{
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
	user := entities.UserEntity{
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