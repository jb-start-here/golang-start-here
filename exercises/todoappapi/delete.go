package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func delete(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := `DELETE FROM todos WHERE id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rw.WriteHeader(http.StatusOK)
}
