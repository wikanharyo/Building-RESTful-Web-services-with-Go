package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Mapping to models with httprouter
	router.ServeFiles("/static/*filepath", http.Dir("/var/www/static"))
	log.Fatal(http.ListenAndServe(":9000", router))
}
