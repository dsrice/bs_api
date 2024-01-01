package repositories

import (
	"app/infra/database/connection"
	"app/repositories/ri"
)

type tokenRepositoryImp struct {
	con *connection.Connection
}

func NewTokenRepository(con *connection.Connection) ri.TokenRepository {
	return &tokenRepositoryImp{con: con}
}