package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello , %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
