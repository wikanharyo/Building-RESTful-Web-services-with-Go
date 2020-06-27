package main

import (
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func myServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world\n")
}

func main() {
	http.HandleFunc("/hello", myServer)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
