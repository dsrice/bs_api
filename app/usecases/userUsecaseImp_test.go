package usecases_test

import (
	"app/controllers/rqp"
	"app/entities/user"
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
	user := user.Entity{LoginID: "t1"}
	s.urepo.On("RegistUser", user).Return(nil)

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	err := uc.RegistUser(&user)

	assert.Nil(s.T(), err)
}

func (s *RegistUserUsecaseSuite) TestFailed() {
	user := user.Entity{LoginID: "t2"}
	s.urepo.On("RegistUser", user).Return(fmt.Errorf("error test"))

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	err := uc.RegistUser(&user)

	assert.Equal(s.T(), err.Error(), "error test")
}

func TestRegistUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(RegistUserUsecaseSuite))
}

type CheckUserUsecaseSuite struct {
	suite.Suite
	urepo *rmock.UserRepositoryMock
}

func (s *CheckUserUsecaseSuite) SetupSuite() {
	s.urepo = new(rmock.UserRepositoryMock)
}

func (s *CheckUserUsecaseSuite) TestSuccess() {
	loginID := "test"
	us := user.Search{
		LoginID: &loginID,
	}
	var ul []*user.Entity
	u := user.Entity{UserID: 1}
	ul = append(ul, &u)
	s.urepo.On("GetUser", &us).Return(ul, nil)

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	user, err := uc.CheckUser(loginID)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), user.UserID, 1)
}

func (s *CheckUserUsecaseSuite) TestSFailedError() {
	loginID := "test2"
	us := user.Search{
		LoginID: &loginID,
	}
	var ul []*user.Entity
	s.urepo.On("GetUser", &us).Return(ul, fmt.Errorf("test error"))

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	user, err := uc.CheckUser(loginID)

	assert.Equal(s.T(), err.Error(), "test error")
	assert.Nil(s.T(), user)
}

func (s *CheckUserUsecaseSuite) TestSFailedUnique() {
	loginID := "test3"
	us := user.Search{
		LoginID: &loginID,
	}
	var ul []*user.Entity
	u := user.Entity{UserID: 1}
	ul = append(ul, &u)
	u = user.Entity{UserID: 2}
	ul = append(ul, &u)
	s.urepo.On("GetUser", &us).Return(ul, nil)

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	user, err := uc.CheckUser(loginID)

	assert.Equal(s.T(), err.Error(), "ユーザーが見つかりませんでいた。 該当数:2")
	assert.Nil(s.T(), user)
}

func TestCheckUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(CheckUserUsecaseSuite))
}

type RegistValidateUsecaseSuite struct {
	suite.Suite
	urepo *rmock.UserRepositoryMock
}

func (s *RegistValidateUsecaseSuite) SetupSuite() {
	s.urepo = new(rmock.UserRepositoryMock)
}

func (s *RegistValidateUsecaseSuite) TestSuccess() {
	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	param := rqp.RegistUser{LoginID: "test", Name: "test2", Password: "testtest"}

	err := uc.RegistValidate(param)

	assert.Nil(s.T(), err)
}

func (s *RegistValidateUsecaseSuite) TestFailedNoLoginID() {
	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	param := rqp.RegistUser{Name: "test2", Password: "testtest"}

	err := uc.RegistValidate(param)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "必須項目が未設定")
}

func (s *RegistValidateUsecaseSuite) TestFailedNoName() {
	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	param := rqp.RegistUser{LoginID: "test2", Password: "testtest"}

	err := uc.RegistValidate(param)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "必須項目が未設定")
}

func (s *RegistValidateUsecaseSuite) TestFailedNoPassword() {
	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	param := rqp.RegistUser{LoginID: "test2", Name: "testtest"}

	err := uc.RegistValidate(param)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "必須項目が未設定")
}

func TestRegistValidateUsecaseSuite(t *testing.T) {
	suite.Run(t, new(RegistValidateUsecaseSuite))
}

type GetUserUsecaseSuite struct {
	suite.Suite
	urepo *rmock.UserRepositoryMock
}

func (s *GetUserUsecaseSuite) SetupSuite() {
	s.urepo = new(rmock.UserRepositoryMock)
}

func (s *GetUserUsecaseSuite) TestSuccess() {
	var ul []*user.Entity
	loginID := "test"
	us := user.Search{
		LoginID: &loginID,
	}

	u := user.Entity{UserID: 1}
	ul = append(ul, &u)
	s.urepo.On("GetUser", &us).Return(ul, nil)

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	result, err := uc.GetUsers(&us)

	assert.Nil(s.T(), err)
	assert.Len(s.T(), result, 1)
}

func (s *GetUserUsecaseSuite) TestFailed() {
	var ul []*user.Entity
	loginID := "t2"
	us := user.Search{
		LoginID: &loginID,
	}

	s.urepo.On("GetUser", &us).Return(ul, fmt.Errorf("test error"))

	ir := ri.InRepository{UserRepo: s.urepo}
	uc := usecases.NewUserUsecase(ir)

	result, err := uc.GetUsers(&us)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "test error")
	assert.Nil(s.T(), result)
}

func TestGetUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(GetUserUsecaseSuite))
}