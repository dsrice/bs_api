package repositories

import (
	"app/entities/user"
	"app/infra/database/connection"
	"app/infra/database/models"
	"app/infra/logger"
	"app/repositories/ri"
	"database/sql"
	"github.com/labstack/echo/v4"
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

func (r *userRepositoryImp) GetUser(us *user.Search, c echo.Context) ([]*user.Entity, error) {
	var muList []qm.QueryMod

	if us.LoginID != nil {
		muList = append(muList, models.UserWhere.LoginID.EQ(*us.LoginID))
	}

	if us.Name != nil {
		muList = append(muList, models.UserWhere.Name.EQ(null.StringFromPtr(us.Name)))
	}

	if us.Mail != nil {
		muList = append(muList, models.UserWhere.Mail.EQ(null.StringFromPtr(us.Mail)))
	}

	ul, err := models.Users(
		muList...,
	).All(c.Request().Context(), r.con)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var uList []*user.Entity
	for _, u := range ul {
		e := user.Entity{}
		e.ConvertUser(u)
		uList = append(uList, &e)
	}

	return uList, nil
}

func (r *userRepositoryImp) RegistUser(user user.Entity, c echo.Context) error {
	logger.Debug("RegistUser Repository start")
	m := user.ConvertUserModel()

	err := m.Insert(c.Request().Context(), r.con, boil.Infer())

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Debug("RegistUser Repository end")
	return nil
}