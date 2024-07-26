package main

import (
	"root/server"
	"testing"
)

// adding one server (2 servers in network)
func TestConnectToNetwork(t *testing.T) {
	go server.New("localhost:5001").Listen()

	Connect("localhost:5000")

	select {}
}
