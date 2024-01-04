package repositories

import (
	"app/entities"
	"app/infra/database/connection"
	"app/infra/database/models"
	"app/infra/logger"
	"app/repositories/ri"
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
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
		logger.Debug("ユーザーが見つかりませんでした")
		return nil, nil
	}

	u := entities.UserEntity{}
	u.ConvertUser(ul[0])

	return &u, nil
}

func (r *userRepositoryImp) RegistUser(user entities.UserEntity) error {
	logger.Debug("RegistUser Repository start")
	m := user.ConvertUserModel()

	err := m.Insert(context.Background(), r.con, boil.Infer())

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Debug("RegistUser Repository end")
	return nil
}