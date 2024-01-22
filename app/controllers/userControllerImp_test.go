package controllers_test

import (
	"app/controllers"
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/controllers/rsp"
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

type UserResitUserControllerSuite struct {
	suite.Suite
	uc *umock.UserUsecaseMock
	ct ci.UserController
}

func (s *UserResitUserControllerSuite) SetupSuite() {
	s.uc = new(umock.UserUsecaseMock)
}

func (s *UserResitUserControllerSuite) TestSuccess() {
	rqp := rqp.RegistUser{LoginID: "t1", Password: "p1", Name: "n1"}
	ue := user.Entity{LoginID: "t1", Password: "p1", Name: "n1"}
	u := new(user.Entity)
	u = nil

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("RegistValidate", rqp).Return(nil)
	s.uc.On("CheckUser", rqp.LoginID, c).Return(u, nil)
	s.uc.On("RegistUser", &ue, c).Return(nil)

	if assert.NoError(s.T(), ct.RegistUser(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)

		result := rsp.User{}
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Equal(s.T(), "t1", result.LoginID)
		assert.Equal(s.T(), "n1", result.Name)
	}
}

func (s *UserResitUserControllerSuite) TestFailLessParam() {
	rqp := rqp.RegistUser{LoginID: "t1"}
	ue := user.Entity{LoginID: "t1", Password: "p1", Name: "n1"}
	u := new(user.Entity)
	u = nil

	rqpJson := `{"login_id": "t1"}`
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(rqpJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("RegistValidate", rqp).Return(fmt.Errorf("error test"))
	s.uc.On("CheckUser", rqp.LoginID, c).Return(u, nil)
	s.uc.On("RegistUser", &ue, c).Return(nil)

	if assert.NoError(s.T(), ct.RegistUser(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}
}

func (s *UserResitUserControllerSuite) TestFailEmptyParam() {
	rqp := rqp.RegistUser{LoginID: "te", Password: "p1", Name: "n1"}
	ue := user.Entity{LoginID: "te", Password: "p1", Name: "n1"}
	u := new(user.Entity)
	u = nil

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("RegistValidate", rqp).Return(fmt.Errorf("error test"))
	s.uc.On("CheckUser", rqp.LoginID, c).Return(u, nil)
	s.uc.On("RegistUser", &ue, c).Return(nil)

	if assert.NoError(s.T(), ct.RegistUser(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}
}

func (s *UserResitUserControllerSuite) TestFailUserLoginID() {
	rqp := rqp.RegistUser{LoginID: "t2", Password: "p2", Name: "n2"}
	user := user.Entity{LoginID: "t2", Password: "p2", Name: "n2"}

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("RegistValidate", rqp).Return(nil)
	s.uc.On("CheckUser", rqp.LoginID, c).Return(&user, nil)
	s.uc.On("RegistUser", &user, c).Return(nil)

	if assert.NoError(s.T(), ct.RegistUser(c)) {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}
}

func (s *UserResitUserControllerSuite) TestFailCheckError() {
	rqp := rqp.RegistUser{LoginID: "t3", Password: "p3", Name: "n3"}
	user := user.Entity{LoginID: "t3", Password: "p3", Name: "n3"}

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("RegistValidate", rqp).Return(nil)
	s.uc.On("CheckUser", rqp.LoginID, c).Return(&user, fmt.Errorf("error test"))
	s.uc.On("RegistUser", &user, c).Return(nil)

	if assert.NoError(s.T(), ct.RegistUser(c)) {
		assert.Equal(s.T(), http.StatusInternalServerError, rec.Code)
	}
}

func (s *UserResitUserControllerSuite) TestFailRegistUser() {
	rqp := rqp.RegistUser{LoginID: "t4", Password: "p4", Name: "n4"}
	ue := user.Entity{LoginID: "t4", Password: "p4", Name: "n4"}
	u := new(user.Entity)
	u = nil

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("RegistValidate", rqp).Return(nil)
	s.uc.On("CheckUser", rqp.LoginID, c).Return(u, nil)
	s.uc.On("RegistUser", &ue, c).Return(fmt.Errorf("error test"))

	if assert.NoError(s.T(), ct.RegistUser(c)) {
		assert.Equal(s.T(), http.StatusInternalServerError, rec.Code)

	}
}

func TestUserResitUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserResitUserControllerSuite))
}

type GetUserUserControllerSuite struct {
	suite.Suite
	uc *umock.UserUsecaseMock
	ct ci.UserController
}

func (s *GetUserUserControllerSuite) SetupTest() {
	s.uc = new(umock.UserUsecaseMock)
}

func (s *GetUserUserControllerSuite) TestSuccess() {
	us := user.Search{}
	u := user.Entity{
		LoginID: "t1",
		Name:    "n1",
	}

	var ul []*user.Entity
	ul = append(ul, &u)

	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("GetUsers", &us, c).Return(ul, nil)

	if assert.NoError(s.T(), ct.GetUser(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)

		var result rsp.GetUser
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Equal(s.T(), "t1", result.Users[0].LoginID)
		assert.Equal(s.T(), "n1", result.Users[0].Name)
	}
}

func (s *GetUserUserControllerSuite) TestFailed() {
	us := user.Search{}

	var ul []*user.Entity

	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	s.uc.On("GetUsers", &us, c).Return(ul, fmt.Errorf("test error"))

	if assert.NoError(s.T(), ct.GetUser(c)) {
		assert.Equal(s.T(), http.StatusInternalServerError, rec.Code)
	}
}

func TestGetUserUserControllerSuite(t *testing.T) {
	suite.Run(t, new(GetUserUserControllerSuite))
}