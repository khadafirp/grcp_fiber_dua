package main

import (
	"grpc_fiber_dua/client"
	"grpc_fiber_dua/server"
)

func main() {
	go server.StartGRPCServer()
	client.StartClient()
}
