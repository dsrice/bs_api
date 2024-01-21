package rsp

import (
	"app/entities/user"
)

type User struct {
	LoginID string `json:"id"`
	Name    string `json:"name"`
	Mail    string `json:"mail"`
}

func (p *User) ConvertResponse(user *user.Entity) {
	p.LoginID = user.LoginID
	p.Name = user.Name
	p.Mail = user.Mail
}

type GetUser struct {
	Users []*User `json:"users"`
}