package controllers_test

import (
	"app/controllers"
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/controllers/rsp"
	"app/entities/token"
	"app/entities/user"
	"app/infra/server"
	"app/usecases/ui"
	"app/usecases/umock"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type LoginControllerSuite struct {
	suite.Suite
	uc *umock.LoginUsecaseMock
	ct ci.LoginController
}

func (s *LoginControllerSuite) SetupSuite() {
	s.uc = new(umock.LoginUsecaseMock)
}

func (s *LoginControllerSuite) TestSuccess() {
	rqp := rqp.Login{LoginID: "t1", Password: "t1"}
	user := user.Entity{UserID: 1, LoginID: "t1"}
	token := token.Entity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)

		result := rsp.Login{}
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Equal(s.T(), "token", result.Token)
		assert.Equal(s.T(), "refresh", result.RefreshToken)
	}
}

func (s *LoginControllerSuite) TestFailBadParam() {
	rqp := rqp.Login{LoginID: "t2"}
	user := user.Entity{UserID: 2, LoginID: "t2"}
	token := token.Entity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login_id:1}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}
}

func (s *LoginControllerSuite) TestFailParam() {
	rqp := rqp.Login{LoginID: "t2"}
	user := user.Entity{UserID: 2, LoginID: "t2"}
	token := token.Entity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login_id":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}
}

func (s *LoginControllerSuite) TestFailNoPass() {
	rqp := rqp.Login{LoginID: "t3", Password: "t3"}
	user := user.Entity{UserID: 3, LoginID: "t3"}
	token := token.Entity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(fmt.Errorf("error test"))
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}
}

func (s *LoginControllerSuite) TestFailNoUser() {
	rqp := rqp.Login{LoginID: "t4", Password: "t4"}
	user := user.Entity{UserID: 4, LoginID: "t4"}
	token := token.Entity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, fmt.Errorf("error test"))
	s.uc.On("GetToken", &user).Return(&token, nil)

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusUnauthorized, rec.Code)
	}
}

func (s *LoginControllerSuite) TestFailToken() {
	rqp := rqp.Login{LoginID: "t5", Password: "t5"}
	user := user.Entity{UserID: 5, LoginID: "t5"}
	token := token.Entity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, fmt.Errorf("error test"))

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusInternalServerError, rec.Code)
	}
}

func TestLoginControllerSuite(t *testing.T) {
	suite.Run(t, new(LoginControllerSuite))
}