package repositories

import (
	"app/entities/token"
	"app/entities/user"
	"app/infra/database/connection"
	"app/infra/database/models"
	"app/infra/logger"
	"app/repositories/ri"
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
)

type tokenRepositoryImp struct {
	con *sql.DB
}

func NewTokenRepository(con *connection.Connection) ri.TokenRepository {
	return &tokenRepositoryImp{con: con.Conn}
}

func (r *tokenRepositoryImp) SetToken(user user.Entity) (*token.Entity, error) {
	logger.Debug("SetToken Start")
	_, err := models.Tokens(models.TokenWhere.UserID.EQ(user.UserID)).DeleteAll(context.Background(), r.con)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	te := token.Entity{User: user}
	te.SetToken()

	token := te.ConvertModel()
	err = token.Insert(context.Background(), r.con, boil.Infer())

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug("SetToken End")
	return &te, nil
}