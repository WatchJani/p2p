package p2p

import (
	"net"
	"root/status"
	"sync"
)

type Peer struct {
	conn chan net.Conn
	peer map[string]net.Conn
	status.Status
	sync.RWMutex
}

func (p Peer) IsNetworkEmpty() bool {
	return len(p.peer) == 0
}

func (p Peer) NetworkSize() int {
	return len(p.conn)
}

func (p *Peer) AddToNetwork(conn net.Conn) {
	p.Lock()
	defer p.Unlock()

	p.peer[conn.RemoteAddr().String()] = conn
}

func (p Peer) IsExist(conn net.Conn) bool {
	p.Lock()
	defer p.Unlock()

	ip := conn.RemoteAddr().String()
	_, ok := p.peer[ip]

	return ok
}

func New(conn chan net.Conn) Peer {
	return Peer{
		conn:   conn,
		peer:   make(map[string]net.Conn),
		Status: status.New(),
	}
}
