package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// CustomServeMux is a struct which can be a multiplexer
type CustomServeMux struct {
}

// Function handler to be overriden
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yor random number is: %f\n", rand.Float64())
}

func main() {
	// Any struct that has serveHTTP function can be a multiplxeer
	mux := &CustomServeMux{}
	http.ListenAndServe(":9000", mux)
}
