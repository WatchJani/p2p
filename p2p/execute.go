package p2p

import (
	"fmt"
	"log"
	"net"
)

// fix all response
func (p *Peer) CanI(conn net.Conn, payload []byte) {
	if p.IsInProcess() {
		p.AddToWaiting(conn)
		WriteMessage(conn, "Info\n{\"info\":\"command on waiting!\"}")
		return
	}

	IpAddr := conn.RemoteAddr().String()
	p.Current = IpAddr //Add mutex

	WriteMessage(conn, "Info\n{\"info\":\"command approve!\"}")

	if p.IsNetworkEmpty() {
		p.AddToNetwork(conn)
		// WriteMessage(conn, "Info\n{\"info\":\"connection added in network!\"}")
		WriteMessage(conn, "AllowToNetwork\n{}")
		p.conn <- conn //add to listen this connection
		return
	}

	msg := fmt.Sprintf("ChangeProcess\n{\"addr\":\"%s\"}", IpAddr)
	if err := p.Broadcast(msg); err != nil {
		log.Println(err)
		p.AddToWaiting(conn)
		WriteMessage(conn, "Info\n{\"info\":\"command on waiting!\"}")
		return
	}
}

// tool to adding new node, is not p2p network just tool
func (p *Peer) AddNode(conn net.Conn, payload []byte) {
	addr, err := ParserIpAddr(payload)
	if err != nil {
		log.Println(err)
		return
	}

	newConn, err := NewConn(addr.Ip)
	if err != nil {
		log.Println(err)
		return
	}
	WriteMessage(newConn, "CanI\n{}")

	//listen on response
	p.conn <- newConn
}

func (p *Peer) AllowToNetwork(conn net.Conn, payload []byte) {
	p.AddToNetwork(conn)
	fmt.Println("Confirm adding")
}

// test information function
func (p *Peer) Info(conn net.Conn, payload []byte) {
	info, err := ParserInfo(payload)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(info.Info)
}

// send to server who in the network
func (p *Peer) ChangeProcess(conn net.Conn, payload []byte) {
	addr, err := ParserIpAddr(payload)
	if err != nil {
		log.Println(err)
		return
	}

	if p.IsInProcess() {
		log.Println("current node is in the process")
		if addr.Ip > p.Current {
			p.Current = addr.Ip
			WriteMessage(conn, "Approve\n{}")
		}

		WriteMessage(conn, "Cancel\n{}")
		return
	}

	p.ChangeProgress()

	p.Current = addr.Ip
	WriteMessage(conn, "Approve\n{}")
}

func (p *Peer) Approve(conn net.Conn, payload []byte) {
	if !p.IsExist(conn) {
		log.Println("This connection is not exist in network")
		return
	}

	if err := p.AddToApprove(conn); err != nil {
		log.Println(err)
		return
	}

	if p.NetworkSize() == p.NumberOfApprovement() {
		//send information for migration
		p.ApproveReset()
	}

	//add some time out
}
