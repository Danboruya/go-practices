package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create new Router
	r := mux.NewRouter()

	// Request handler
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// Second parameter of http.ListenAndServe is main router of HTTP server.
	// When set nil, this function use default router of net/http package.
	http.ListenAndServe(":80", r)
}