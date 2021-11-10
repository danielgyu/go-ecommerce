package product

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
		Addr:   "golangdb:3306",
		DBName: "ecommerce",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
