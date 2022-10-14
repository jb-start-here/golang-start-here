package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	var todos []todo
	rw.Header().Set("Content-Type", "application/json")

	rows, err := db.Query(`SELECT * FROM todos`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo todo
		if err := rows.Scan(&todo.id, &todo.description, &todo.done, &todo.duedate); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	res, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(rw, string(res))
}
