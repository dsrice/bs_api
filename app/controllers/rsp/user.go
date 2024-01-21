package rsp

import (
	"app/entities/user"
)

type User struct {
	LoginID string
	Name    string
	Mail    string
}

func (p *User) ConvertResponse(user *user.Entity) {
	p.LoginID = user.LoginID
	p.Name = user.Name
	p.Mail = user.Mail
}