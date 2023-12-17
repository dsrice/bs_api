package connection

import (
	"app/infra/logger"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Connection struct {
	Conn *sql.DB
}

func (c *Connection) GetConnection() error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	conf := mysql.Config{
		DBName:    "bowling_score",
		User:      "user",
		Passwd:    "password",
		Addr:      "db:3306",
		Net:       "tcp",
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		logger.Error(err.Error())
		return err
	} else {
		c.Conn = db
	}

	return nil
}