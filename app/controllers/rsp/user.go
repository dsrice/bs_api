package rsp

import "app/entities"

type User struct {
	LoginID string
	Name    string
	Mail    string
}

func (p *User) ConvertResponse(user *entities.UserEntity) {
	p.LoginID = user.LoginID
	p.Name = user.Name
	p.Mail = user.Mail
}