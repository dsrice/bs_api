package tester

import (
	"app/infra/server"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/http/httptest"
)

type Tester struct {
	Request  *http.Request
	Recorder *httptest.ResponseRecorder
	Contest  echo.Context
}

func CreateContext(method, path string, body io.Reader) Tester {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}

	return Tester{
		Request:  req,
		Recorder: rec,
		Contest:  e.NewContext(req, rec),
	}
}