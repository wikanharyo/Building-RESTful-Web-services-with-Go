package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func main() {
	// Create a web server using restful.Webservice struct
	webservice := new(restful.WebService)
	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// Add the service to application (register webservice with the go-restful)
	restful.Add(webservice)
	http.ListenAndServe(":9000", nil)
}

func pingTime(req *restful.Request, resp *restful.Response) {
	// Write ti the response
	io.WriteString(resp, fmt.Sprintf("%s\n", time.Now()))
}
