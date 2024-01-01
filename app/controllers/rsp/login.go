package rsp

import "app/entities"

type Login struct {
	Token        string
	RefreshToken string
}

func (p *Login) ConvertResponse(token *entities.TokenEntity) {
	p.Token = token.Token
	p.RefreshToken = token.RefreshToken
}