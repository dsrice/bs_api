package user

import (
	"app/infra/database/models"
	"github.com/volatiletech/null"
)

type Entity struct {
	UserID   int
	LoginID  string
	Name     string
	Mail     string
	Password string
}

func (e *Entity) ConvertUser(m *models.User) {
	e.UserID = m.ID
	e.LoginID = m.LoginID
	e.Name = m.Name.String
	e.Mail = m.Mail.String
}

func (e *Entity) ConvertUserModel() models.User {
	return models.User{
		ID:       e.UserID,
		LoginID:  e.LoginID,
		Name:     null.String{String: e.Name, Valid: true},
		Password: e.Password,
		Mail:     null.String{String: e.Mail, Valid: len(e.Mail) > 0},
	}
}