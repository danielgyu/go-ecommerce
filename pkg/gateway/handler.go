package gateway

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type gatewayHandler struct {
	clients *grpcClients
}

func NewGatewayHandler(c *grpcClients) *gatewayHandler {
	return &gatewayHandler{clients: c}
}

func errorResponse(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(400)
	w.Write([]byte("bad request"))
}

func (h *gatewayHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("healthy"))
}

func (h *gatewayHandler) initializedb(w http.ResponseWriter, r *http.Request) {
	cfg := mysql.Config{
		User:                 "test",
		Passwd:               "test",
		Net:                  "tcp",
		Addr:                 "golangdb:3306",
		DBName:               "ecommerce",
		AllowNativePasswords: true,
		MultiStatements:      true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println("error opening mysql")
		log.Fatal(err)
	}

	init_script, iOerr := ioutil.ReadFile("./initdb.sql")
	if iOerr != nil {
		log.Println("error reading sql")
		log.Fatal(err)
	}

	_, err = db.Exec(string(init_script))
	if err != nil {
		log.Println("error executing sql statement")
		log.Fatal(err)
	}

	w.Write([]byte("db initialization success"))
}
