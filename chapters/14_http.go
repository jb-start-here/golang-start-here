package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Go lang provides great http utilities as a part of stdlib. its called net package. One of the sub package is called http
// This http sub package provides decent abstarctions as a http request client (GET, HEAD, POST) and also creating http servers (session and cookie management),
// file servers and routing etc

// Some basic contructs in this package are

// A Request represents an HTTP request received by a server, sent by a client.
// type Request struct

// Response represents the response from an HTTP request.
// type Response struct

// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// type ResponseWriter interface

func main() {
	// makeHTTPRequests()
	// basicHTTPServer()
	// basicHTTPServerHandlerFunc()
	// anotherSampleServer()
	routedServer()
}

func makeHTTPRequests() {
	var googleResponse *http.Response
	var err error
	googleResponse, err = http.Get("https://www.google.com")

	if err != nil {
		log.Println(err)
	}
	fmt.Println(googleResponse.Status)

	// Head and Post are also available
}

// A Handler responds to an HTTP request. The ServeHTTP method that this interface requires
// writes reply headers and data to the ResponseWriter and then returns.
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// Lets create something that can satisfy the interface of a Handler
type greeter struct {
	name    string
	counter int
}

func (g greeter) greetingMsg() string {
	return fmt.Sprintf("Hello %s, how are you?. counter -> %d", g.name, g.counter)
}

// This satisfied the Handler interface
func (g *greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// r has all the info about the request

	if r.Method == "GET" {
		g.counter++
		fmt.Fprintf(w, g.greetingMsg())
	} else {
		http.Error(w, "Method not allowed", 405)
	}
}

func basicHTTPServer() {
	myGreeter := &greeter{
		"Sherlock Holmes",
		0,
	}

	// http has `Handle` method that takes a path and handler and invokes handler when path matched
	http.Handle("/greet", myGreeter)

	// The ListenAndServe listens on the TCP network address and then handles requests on incoming connections.
	// The second arg to the listen funciton is an optional router - well see more later

	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// The HandlerFunc type is an adapter that allows the use of ordinary functions as HTTP handlers.
// type HandlerFunc func(ResponseWriter, *Request)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func basicHTTPServerHandlerFunc() {
	http.HandleFunc("/hello", helloWorldHandler)

	// can also use anon functions
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Goodbye World!")
	})

	log.Fatal(http.ListenAndServe(":5050", nil))
}

func anotherSampleServer() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// also lets send a status code along with a text
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World!")
	})

	// for 404
	http.HandleFunc("/gimme404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 - not found")
	})

	// for getting request headers
	http.HandleFunc("/ua", func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		fmt.Fprintf(w, "User agent: %s\n", ua)
	})

	// Queries in urls
	http.HandleFunc("/queries", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["name"]
		if ok {
			fmt.Fprintf(w, "Hello %s!\n", keys[0])
		} else {
			fmt.Fprintf(w, "Hello %s!\n", "Guest")
		}
	})

	log.Fatal(http.ListenAndServe(":5050", nil))
}

// More info here https://zetcode.com/golang/http-server/

// Sometimes we need more complex routing so instead we can define a custom router and attach it to the listen and serve call
func routedServer() {
	// The NewServeMux function allocates and returns a new ServeMux.

	// ServeMux is an HTTP request multiplexer. It is used for request routing and dispatching.
	// The request routing is based on URL patterns. Each incoming request's URL is matched
	// against a list of registered patterns. A handler for the pattern that most closely fits the URL is called

	mux := http.NewServeMux()

	mux.Handle("/greet", &greeter{
		"Sherlock Holmes",
		0,
	})
	// mux also takes a Handler interface i.e it must have method serveHTTP

	// or we could also do use handlerfunc adapter

	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, time.Now().String())
	})

	http.ListenAndServe(":5050", mux)

	// in the past examples we passed in nil as a the second argument
	// DefaultServeMux is a just a ServeMux. It is used when we pass nil to the ListenAndServe method for the second parameter.
}

// The standard multiplexer is limited in its funcitonality so devs tend to use third part go mux libraries
// which are still compatible with the Handler interface therefore we can pass it to ListenAndServe
// example; https://github.com/gorilla/mux

func fileServer() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.ListenAndServe(":5050", nil)
}
