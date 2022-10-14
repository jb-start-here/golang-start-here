package main

import (
	"encoding/json"
	"time"
)

type todo struct {
	id          int
	description string
	done        bool
	duedate     time.Time
}

// We have to satisfy the json.Marshaler interface which needs the MarshalJSON method
func (t todo) MarshalJSON() ([]byte, error) {
	todoReplica := struct {
		Id          int
		Description string
		Done        bool
		Duedate     string
	}{
		Id:          t.id,
		Description: t.description,
		Done:        t.done,
		Duedate:     t.duedate.Format(time.RFC3339),
	}

	res, err := json.Marshal(todoReplica)
	if err != nil {
		return nil, err
	}
	return res, nil
}
