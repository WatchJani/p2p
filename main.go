package main

import (
	"fmt"
	"log"
	"root/p2p"
	"root/server"
)

func main() {
	server.New("localhost:5000").Listen()
}

func Connect(addr string) {
	conn, err := p2p.NewConn("localhost:5000")
	if err != nil {
		log.Println(err)
		return
	}

	msg := fmt.Sprintf("AddNode\n{\"addr\":\"%s\"}", addr)
	p2p.WriteMessage(conn, msg)
}
