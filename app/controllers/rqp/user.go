package rqp

import "app/entities"

type RegistUser struct {
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Mail     string `json:"mail"`
}

func (e *RegistUser) ConvertEntity() entities.UserEntity {
	return entities.UserEntity{
		LoginID:  e.LoginID,
		Password: e.Password,
		Mail:     e.Mail,
		Name:     e.Name,
	}
}