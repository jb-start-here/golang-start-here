package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "gopwd"
	dbname   = "postgres"
)

var db *sql.DB

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(_ http.ResponseWriter, _ *http.Request) {})
	router.HandleFunc("/todos", index).Methods("GET")
	router.HandleFunc("/todos/{id}", show).Methods("GET")
	router.HandleFunc("/todos/{id}", delete).Methods("DELETE")

	initDBConn()
	listenAndServe(router)
}

func initDBConn() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	if db, err = sql.Open("postgres", psqlconn); err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func listenAndServe(router *mux.Router) {
	if err := http.ListenAndServe(":5050", router); err != nil {
		log.Fatal(err)
	}
}
