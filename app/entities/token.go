package entities

import (
	"app/entities/user"
	"app/infra/database/models"
	"github.com/volatiletech/null"
	"math/rand"
)

type TokenEntity struct {
	User         user.Entity
	Token        string
	RefreshToken string
}

func (e *TokenEntity) ConvertModel() *models.Token {
	return &models.Token{
		UserID:       e.User.UserID,
		Token:        null.String{String: e.Token, Valid: true},
		RefreshToken: null.String{String: e.RefreshToken, Valid: true},
	}
}

func (e *TokenEntity) SetToken() {
	e.Token = e.createToken()
	e.RefreshToken = e.createToken()
}

func (e *TokenEntity) createToken() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 30)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}