package rsp

import (
	"app/entities/token"
)

type Login struct {
	Token        string
	RefreshToken string
}

func (p *Login) ConvertResponse(token *token.Entity) {
	p.Token = token.Token
	p.RefreshToken = token.RefreshToken
}