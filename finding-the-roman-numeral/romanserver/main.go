package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/wikanharyo/finding-the-roman-numeral/romanNumerals"
)

func main() {
	// http package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// url path for the HTTP request and split the path
		urlPathElements := strings.Split(r.URL.Path, "/")
		// if request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number < 1 || number > 10 {
				// if resource is not in the list, send Not Found status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			// for all other request, tell that client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
