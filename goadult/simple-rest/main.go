package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if len(name) == 0 {
		name = "Stranger"
	}
	fmt.Fprintf(w, "Hello, %v", name)
}

// GET http://localhost:8080/hello will return

// Hello, Stranger
// and GET http://localhost:8080/hello?name=Jack will result in

// Hello, Jack
