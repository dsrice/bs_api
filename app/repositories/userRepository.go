package repositories

import (
	"app/infra/database/connection"
	"app/repositories/ri"
)

type userRepositoryImp struct {
	con *connection.Connection
}

func NewUserRepository(con *connection.Connection) ri.UserRepository {
	return &userRepositoryImp{con}
}