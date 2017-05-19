package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/foo", Foo)

	http.Handle("/", r)
	http.Handle("/foo", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Hello greets the world
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

// Foo is greeting only himself
func Foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello fooo!")
}
