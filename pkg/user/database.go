package user

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlClient() *sql.DB {
	cfg := mysql.Config{
		User:   "test",
		Passwd: "test",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "ecommerce",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
