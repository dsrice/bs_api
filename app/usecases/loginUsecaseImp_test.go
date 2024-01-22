package usecases_test

import (
	"app/controllers/rqp"
	"app/entities/token"
	"app/entities/user"
	"app/repositories/ri"
	"app/repositories/rmock"
	"app/tester"
	"app/usecases"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
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
	urepo  *rmock.UserRepositoryMock
	tester tester.Tester
}

func (s *GetUserLoginUsecaseSuite) SetupSuite() {
	s.urepo = new(rmock.UserRepositoryMock)
	s.tester = tester.CreateContext(http.MethodPost, "/", nil)
}

func (s *GetUserLoginUsecaseSuite) TestSuccess() {
	ue := user.Entity{UserID: 1, LoginID: "t1"}
	loginID := "t1"
	var ul []*user.Entity
	ul = append(ul, &ue)
	us := user.Search{
		LoginID: &loginID,
	}
	s.urepo.On("GetUser", &us, s.tester.Context).Return(ul, nil)

	repo := ri.InRepository{UserRepo: s.urepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetUser("t1", s.tester.Context)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result.UserID, 1)
	assert.Equal(s.T(), result.LoginID, "t1")
}

func (s *GetUserLoginUsecaseSuite) TestError() {
	ue := user.Entity{UserID: 1, LoginID: "t2"}
	loginID := "t2"
	us := user.Search{
		LoginID: &loginID,
	}
	var ul []*user.Entity
	ul = append(ul, &ue)

	s.urepo.On("GetUser", &us, s.tester.Context).Return(ul, fmt.Errorf("error test"))

	repo := ri.InRepository{UserRepo: s.urepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetUser("t2", s.tester.Context)

	assert.Nil(s.T(), result)
	assert.Equal(s.T(), err.Error(), "error test")
}

func TestGetUserLoginUsecaseSuite(t *testing.T) {
	suite.Run(t, new(GetUserLoginUsecaseSuite))
}

type GetTokenLoginUsecaseSuite struct {
	suite.Suite
	trepo  *rmock.TokenRepositoryMock
	tester tester.Tester
}

func (s *GetTokenLoginUsecaseSuite) SetupSuite() {
	s.trepo = new(rmock.TokenRepositoryMock)
	s.tester = tester.CreateContext(http.MethodPost, "/", nil)
}

func (s *GetTokenLoginUsecaseSuite) TestSuccess() {
	user := user.Entity{UserID: 1, LoginID: "t1"}
	token := token.Entity{User: user}
	s.trepo.On("SetToken", user, s.tester.Context).Return(&token, nil)

	repo := ri.InRepository{TokenRepository: s.trepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetToken(&user, s.tester.Context)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), result.User, user)
}

func (s *GetTokenLoginUsecaseSuite) TestFail() {
	user := user.Entity{UserID: 2, LoginID: "t2"}
	token := token.Entity{User: user}
	s.trepo.On("SetToken", user, s.tester.Context).Return(&token, fmt.Errorf("error test"))

	repo := ri.InRepository{TokenRepository: s.trepo}

	uc := usecases.NewLoginUsecase(repo)

	result, err := uc.GetToken(&user, s.tester.Context)

	assert.Nil(s.T(), result)
	assert.Equal(s.T(), err.Error(), "error test")
}

func TestGetTokenLoginUsecaseSuite(t *testing.T) {
	suite.Run(t, new(GetTokenLoginUsecaseSuite))
}