package repository

import (
	"database/sql"
	"log"
	"os"
)

func GetDbPool() *sql.DB {
	connSource := os.Getenv("MYSQL_CONNECTION_STRING")

	db, err := sql.Open("mysql", connSource)
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(4)
	return db
}
