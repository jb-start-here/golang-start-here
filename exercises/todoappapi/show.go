package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func show(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var todo todo
	query := `SELECT * FROM todos WHERE id = $1`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&todo.id, &todo.description, &todo.done, &todo.duedate); err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintln(rw, "{}")
			return
		} else {
			log.Fatal(err)
		}
	}

	res, err := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(rw, string(res))
}
