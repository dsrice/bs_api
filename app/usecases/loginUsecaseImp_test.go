package usecases_test

import (
	"app/controllers/rqp"
	"app/entities"
	"app/repositories/ri"
	"app/repositories/rmock"
	"app/usecases"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidateLoginUsecaseSuite struct {
	suite.Suite
	urepo *rmock.UserRepositoryMock
}

func (s *ValidateLoginUsecaseSuite) SetupSuite() {
	s.urepo = new(rmock.UserRepositoryMock)
}

func TestValidateLoginUsecaseSuite(t *testing.T) {
	suite.Run(t, new(ValidateLoginUsecaseSuite))
}

func (s *ValidateLoginUsecaseSuite) TestSuccess() {
	repo := ri.InRepository{UserRepo: s.urepo}

	uc := usecases.NewLoginUsecase(repo)

	param := rqp.Login{LoginID: "t1", Password: "t2"}
	err := uc.Validate(param)

	assert.Nil(s.T(), err)
}

func (s *ValidateLoginUsecaseSuite) TestFailNoLoginID() {
	repo := ri.InRepository{UserRepo: s.urepo}

	uc := usecases.NewLoginUsecase(repo)

	param := rqp.Login{LoginID: "", Password: "t2"}
	err := uc.Validate(param)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "ログインIDもしくはパスワードがありません")
}

func (s *ValidateLoginUsecaseSuite) TestFailNoPassword() {
	repo := ri.InRepository{UserRepo: s.urepo}

	uc := usecases.NewLoginUsecase(repo)

	param := rqp.Login{LoginID: "t1", Password: ""}
	err := uc.Validate(param)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "ログインIDもしくはパスワードがありません")
}

type GetUserLoginUsecaseSuite struct {
	suite.Suite
	urepo *rmock.UserRepositoryMock
}

func (s *GetUserLoginUsecaseSuite) SetupSuite() {
	s.urepo = new(rmock.UserRepositoryMock)
}

func (s *GetUserLoginUsecaseSuite) TestSuccess() {
	user := entities.UserEntity{UserID: 1, LoginID: "t1"}
	s.urepo.On("GetUser", "t1").Return(&user, nil)

	repo := ri.InRepository{UserRepo: s.urepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetUser("t1")

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result.UserID, 1)
	assert.Equal(s.T(), result.LoginID, "t1")
}

func (s *GetUserLoginUsecaseSuite) TestError() {
	user := entities.UserEntity{UserID: 1, LoginID: "t1"}
	s.urepo.On("GetUser", "t2").Return(&user, fmt.Errorf("error test"))

	repo := ri.InRepository{UserRepo: s.urepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetUser("t2")

	assert.Nil(s.T(), result)
	assert.Equal(s.T(), err.Error(), "error test")
}

func TestGetUserLoginUsecaseSuite(t *testing.T) {
	suite.Run(t, new(GetUserLoginUsecaseSuite))
}

type GetTokenLoginUsecaseSuite struct {
	suite.Suite
	trepo *rmock.TokenRepositoryMock
}

func (s *GetTokenLoginUsecaseSuite) SetupSuite() {
	s.trepo = new(rmock.TokenRepositoryMock)
}

func (s *GetTokenLoginUsecaseSuite) TestSuccess() {
	user := entities.UserEntity{UserID: 1, LoginID: "t1"}
	token := entities.TokenEntity{User: user}
	s.trepo.On("SetToken", user).Return(&token, nil)

	repo := ri.InRepository{TokenRepository: s.trepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetToken(&user)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result.User, user)
}

func (s *GetTokenLoginUsecaseSuite) TestFail() {
	user := entities.UserEntity{UserID: 2, LoginID: "t2"}
	token := entities.TokenEntity{User: user}
	s.trepo.On("SetToken", user).Return(&token, fmt.Errorf("error test"))

	repo := ri.InRepository{TokenRepository: s.trepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetToken(&user)

	assert.Nil(s.T(), result)
	assert.Equal(s.T(), err.Error(), "error test")
}

func TestGetTokenLoginUsecaseSuite(t *testing.T) {
	suite.Run(t, new(GetTokenLoginUsecaseSuite))
}