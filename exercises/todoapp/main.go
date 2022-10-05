package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitedUrl := r.URL.Path
		method := r.Method
		log.Println(method + " " + visitedUrl + " " + time.Now().String())
		fmt.Fprintf(w, "Welcome to new server! You have visited: " + visitedUrl)
	})
	http.HandleFunc("/favicon.ico", http.NotFound)

	// listen to port
	err := http.ListenAndServe(":5050", nil)
	log.Fatal(err)
}
