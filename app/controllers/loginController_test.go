package controllers_test

import (
	"app/controllers"
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/controllers/rsp"
	"app/entities"
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

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (s *LoginControllerSuite) SetupSuite() {
	s.uc = new(umock.LoginUsecaseMock)
}

func (s *LoginControllerSuite) TestSuccess() {
	rqp := rqp.Login{LoginID: "t1", Password: "t1"}
	user := entities.UserEntity{UserID: 1, LoginID: "t1"}
	token := entities.TokenEntity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
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

	s.uc = new(umock.LoginUsecaseMock)
}

func (s *LoginControllerSuite) TestFailBadParam() {
	rqp := rqp.Login{LoginID: "t2"}
	user := entities.UserEntity{UserID: 2, LoginID: "t2"}
	token := entities.TokenEntity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login_id:1}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}

	s.uc = new(umock.LoginUsecaseMock)
}

func (s *LoginControllerSuite) TestFailParam() {
	rqp := rqp.Login{LoginID: "t2"}
	user := entities.UserEntity{UserID: 2, LoginID: "t2"}
	token := entities.TokenEntity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login_id":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}

	s.uc = new(umock.LoginUsecaseMock)
}

func (s *LoginControllerSuite) TestFailNoPass() {
	rqp := rqp.Login{LoginID: "t3", Password: "t3"}
	user := entities.UserEntity{UserID: 3, LoginID: "t3"}
	token := entities.TokenEntity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(fmt.Errorf("error test"))
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, nil)

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}

	s.uc = new(umock.LoginUsecaseMock)
}

func (s *LoginControllerSuite) TestFailNoUser() {
	rqp := rqp.Login{LoginID: "t4", Password: "t4"}
	user := entities.UserEntity{UserID: 4, LoginID: "t4"}
	token := entities.TokenEntity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, fmt.Errorf("error test"))
	s.uc.On("GetToken", &user).Return(&token, nil)

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusUnauthorized, rec.Code)
	}

	s.uc = new(umock.LoginUsecaseMock)
}

func (s *LoginControllerSuite) TestFailToken() {
	rqp := rqp.Login{LoginID: "t5", Password: "t5"}
	user := entities.UserEntity{UserID: 5, LoginID: "t5"}
	token := entities.TokenEntity{User: user, Token: "token", RefreshToken: "refresh"}
	s.uc.On("Validate", rqp).Return(nil)
	s.uc.On("GetUser", rqp.LoginID).Return(&user, nil)
	s.uc.On("GetToken", &user).Return(&token, fmt.Errorf("error test"))

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{Login: s.uc}
	ct := controllers.NewLoginController(ic)

	if assert.NoError(s.T(), ct.Login(c)) {
		assert.Equal(s.T(), http.StatusInternalServerError, rec.Code)
	}

	s.uc = new(umock.LoginUsecaseMock)
}

func TestLoginControllerSuite(t *testing.T) {
	suite.Run(t, new(LoginControllerSuite))
}