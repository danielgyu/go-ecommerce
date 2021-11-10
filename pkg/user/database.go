package user

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlClient() *sql.DB {
	cfg := mysql.Config{
		User:                 "test",
		Passwd:               "test",
		Net:                  "tcp",
		Addr:                 "golangdb:3306",
		AllowNativePasswords: true,
		DBName:               "ecommerce",
	}

	log.Printf("connecting to database at %s", cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connection complete")

	return db
}
