package repositories

import (
	"app/entities"
	"app/infra/database/connection"
	"app/infra/database/models"
	"app/infra/logger"
	"app/repositories/ri"
	"context"
	"database/sql"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type userRepositoryImp struct {
	con *sql.DB
}

func NewUserRepository(con *connection.Connection) ri.UserRepository {
	return &userRepositoryImp{con.Conn}
}

func (r *userRepositoryImp) GetUser(loginID, name, mail *string) ([]*entities.UserEntity, error) {
	var muList []qm.QueryMod

	if loginID != nil {
		muList = append(muList, models.UserWhere.LoginID.EQ(*loginID))
	}

	if name != nil {
		muList = append(muList, models.UserWhere.Name.EQ(null.StringFromPtr(name)))
	}

	if mail != nil {
		muList = append(muList, models.UserWhere.Mail.EQ(null.StringFromPtr(mail)))
	}

	ul, err := models.Users(
		muList...,
	).All(context.Background(), r.con)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var uList []*entities.UserEntity
	for _, u := range ul {
		e := entities.UserEntity{}
		e.ConvertUser(u)
		uList = append(uList, &e)
	}

	return uList, nil
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