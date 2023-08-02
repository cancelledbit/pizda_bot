package repository

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	"os"
)

func GetDbPool() *sql.DB {
	connSource := os.Getenv("MYSQL_CONNECTION_STRING")

	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	db := sqldblogger.OpenDriver(connSource,
		&mysql.MySQLDriver{},
		loggerAdapter,
	)

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(4)
	return db
}
