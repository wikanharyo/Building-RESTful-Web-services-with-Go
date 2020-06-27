package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// HOlds information about arguments passed from RPC Client to the server
type Args struct{}

// Create TimeServer to register with the rpc, server w
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// Create a new RPC server
	timeserver := new(TimeServer)
	// Register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// Listen for requests on port
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)

}
