package server

import (
	"app/entities/user"
	"app/infra/logger"
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetSession(c echo.Context) *sessions.Session {
	s, err := session.Get("session", c)
	gob.Register(user.Entity{})

	if err != nil {
		logger.Error(err.Error())
	}

	return s
}