package server

import (
	"app/controllers/ci"
	"app/infra/logger"
	"app/repositories/ri"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	echo  *echo.Echo
	Login ci.LoginController
	User  ci.UserController
	TRepo ri.TokenRepository
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewServer(s ci.InController, tr ri.TokenRepository) *Server {
	return &Server{
		Login: s.Login,
		User:  s.UserController,
		TRepo: tr,
	}
}

func (s *Server) Start() {
	s.echo = echo.New()

	//セッションを設定
	store := sessions.NewCookieStore([]byte("bowring_api_secret_key"))
	store.Options.Path = "/"
	store.Options.MaxAge = -1
	store.Options.HttpOnly = true
	s.echo.Use(session.Middleware(store))

	s.echo.Validator = &CustomValidator{Validator: validator.New()}

	ka := middleware.KeyAuthConfig{
		Skipper: func(c echo.Context) bool {
			logger.Debug(c.Request().URL.Path)
			if c.Request().URL.Path == "/login" {
				return true
			}

			return false
		},
		KeyLookup:  "header:" + echo.HeaderAuthorization,
		AuthScheme: "Bearer",
		Validator: func(auth string, c echo.Context) (bool, error) {
			logger.Debug("Auth Start")
			logger.Debug("token:" + auth)

			sess := GetSession(c)

			u, err := s.TRepo.SearchUser(auth, c)

			if err != nil {
				logger.Error(err.Error())
				return false, err
			}

			sess.Values["user"] = u

			err = sess.Save(c.Request(), c.Response())

			if err != nil {
				logger.Error(err.Error())
				return true, err
			}

			logger.Debug("Auth End")
			return true, nil
		},
	}

	s.echo.Use(middleware.KeyAuthWithConfig(ka))
	s.routing()

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}