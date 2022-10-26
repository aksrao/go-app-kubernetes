package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome,%q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are in home page")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}