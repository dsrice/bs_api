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

func NewConnection() *Connection {
	c, err := getConnection()

	if err != nil {
		panic(err)
	}

	return &Connection{Conn: c}
}

func getConnection() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	conf := mysql.Config{
		DBName:               "bowling_score",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "db:3306",
		Net:                  "tcp",
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	conn, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return conn, nil
}