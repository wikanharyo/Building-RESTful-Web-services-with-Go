package main

import (
	"fmt"
	"net/http"
)

// Requset==>Middleware==>MainHandler

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase1")
		// Pass control back to the handler
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after respons phase!")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("Ok"))
}
func main() {
	// HandlerFunc returns a HTTP Handler
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":9000", nil)
}
