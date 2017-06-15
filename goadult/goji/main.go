// 1. Register function hello to handle requests for path /hello
// 2. Create and run a web server listening on a specified port
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

type book struct {
	ISBN    string "json:isbn"
	Title   string "json:name"
	Authors string "json:author"
	Price   string "json:price"
}

var bookStore = []book{
	book{
		ISBN:    "0321774639",
		Title:   "Programming in Go: Creating Applications for the 21st Century (Developer's Library)",
		Authors: "Mark Summerfield",
		Price:   "$34.57",
	},
	book{
		ISBN:    "0134190440",
		Title:   "The Go Programming Language",
		Authors: "Alan A. A. Donovan, Brian W. Kernighan",
		Price:   "$34.57",
	},
}

func main() {
	mux := goji.NewMux()
	// Registering handlers in Goji is very similar to net/http
	mux.HandleFunc(pat.Get("/books"), allBooks)
	mux.HandleFunc(pat.Get("/books/:isbn"), bookByISBN)
	// Registering middleware
	mux.Use(logging)
	http.ListenAndServe("localhost:8080", mux)
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	jsonOut, _ := json.Marshal(bookStore)
	fmt.Fprintf(w, string(jsonOut))
}

func bookByISBN(w http.ResponseWriter, r *http.Request) {
	isbn := pat.Param(r, "isbn")
	for _, b := range bookStore {
		if b.ISBN == isbn {
			jsonOut, _ := json.Marshal(b)
			fmt.Fprintf(w, string(jsonOut))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// logging middleware
// prints requested URL and relays to handler for further processing
func logging(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %v\n", r.URL)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// [ `go run main.go` | done: 722.5562ms ]
