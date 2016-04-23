package main

import (
	"fmt"
	"log"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)


func main2() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index2)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func Index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}