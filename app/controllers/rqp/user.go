package rqp

import (
	"app/entities/user"
)

type RegistUser struct {
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Mail     string `json:"mail"`
}

func (e *RegistUser) ConvertEntity() user.Entity {
	return user.Entity{
		LoginID:  e.LoginID,
		Password: e.Password,
		Mail:     e.Mail,
		Name:     e.Name,
	}
}