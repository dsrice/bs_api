package entities

import "app/infra/database/models"

type UserEntity struct {
	UserID  int
	LoginID string
}

func (e *UserEntity) ConvertUser(m *models.User) {
	e.UserID = m.ID
	e.LoginID = m.LoginID
}