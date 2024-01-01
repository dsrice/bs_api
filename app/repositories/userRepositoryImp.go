package repositories

import (
	"app/entities"
	"app/infra/database/connection"
	"app/infra/database/models"
	"app/infra/logger"
	"app/repositories/ri"
	"context"
	"database/sql"
	"fmt"
)

type userRepositoryImp struct {
	con *sql.DB
}

func NewUserRepository(con *connection.Connection) ri.UserRepository {
	return &userRepositoryImp{con.Conn}
}

func (r *userRepositoryImp) GetUser(loginID string) (*entities.UserEntity, error) {
	ul, err := models.Users(
		models.UserWhere.LoginID.EQ(loginID),
	).All(context.Background(), r.con)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(ul) != 1 {
		err = fmt.Errorf("対象ユーザが見つかりませんでした")
		logger.Error(err.Error())
		return nil, err
	}

	u := entities.UserEntity{}
	u.ConvertUser(ul[0])

	return &u, nil
}