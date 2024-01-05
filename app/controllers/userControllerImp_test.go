package controllers_test

import (
	"app/controllers"
	"app/controllers/ci"
	"app/controllers/rqp"
	"app/controllers/rsp"
	"app/entities"
	"app/infra/server"
	"app/usecases/ui"
	"app/usecases/umock"
	"encoding/json"
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
	user := entities.UserEntity{LoginID: "t1", Password: "p1", Name: "n1"}
	u := new(entities.UserEntity)
	u = nil
	s.uc.On("RegistValidate", rqp).Return(nil)
	s.uc.On("CheckUser", rqp.LoginID).Return(u, nil)
	s.uc.On("RegistUser", &user).Return(nil)

	rqpJson, _ := json.Marshal(rqp)
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(rqpJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ic := ui.InUsecase{UserUsecase: s.uc}
	ct := controllers.NewUserController(ic)

	if assert.NoError(s.T(), ct.RegistUser(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)

		result := rsp.User{}
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Equal(s.T(), "t1", result.LoginID)
		assert.Equal(s.T(), "n1", result.Name)
	}
}

func TestUserResitUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserResitUserControllerSuite))
}