package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Function handler ArticleHandler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns all path parameters as map
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	// Create a new router
	r := mux.NewRouter()
	// Attach a path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	// Reverse mapping URL, HATI2 HANDLERFUNC!=HANDLEFUNC
	// r.HandlerFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")
	// url, _ := r.Get("articleRoute").URL("category", "books", "id", "123")
	// fmt.Printf(url.Path) // prints /articles/books/123
	// ....

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:9000",
		// Enforce timouts for severs
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
