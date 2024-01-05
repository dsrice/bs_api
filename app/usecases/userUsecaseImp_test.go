package usecases_test

import (
	"app/entities"
	"app/repositories/ri"
	"app/repositories/rmock"
	"app/usecases"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RegistUserUsecaseSuite struct {
	suite.Suite
	urepo *rmock.UserRepositoryMock
}

func (s *RegistUserUsecaseSuite) SetupSuite() {
	s.urepo = new(rmock.UserRepositoryMock)
}

func (s *RegistUserUsecaseSuite) TestSuccess() {
	user := entities.UserEntity{LoginID: "t1"}
	s.urepo.On("RegistUser", user).Return(nil)

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	err := uc.RegistUser(&user)

	assert.Nil(s.T(), err)
}

func (s *RegistUserUsecaseSuite) TestFailed() {
	user := entities.UserEntity{LoginID: "t2"}
	s.urepo.On("RegistUser", user).Return(fmt.Errorf("error test"))

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	err := uc.RegistUser(&user)

	assert.Equal(s.T(), err.Error(), "error test")
}

func TestVRegistUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(RegistUserUsecaseSuite))
}