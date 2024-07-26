package router

import (
	"net"
	"root/p2p"
	"sync"
)

type HandlerFunc func(conn net.Conn, payload []byte)

type Router struct {
	path  map[string]HandlerFunc
	Parse func([]byte) (string, []byte)
	sync.RWMutex
}

func New(p2p p2p.Peer) Router {
	return Router{
		path:  RouterP2P(p2p),
		Parse: ParserJSON,
	}
}

func (c *Router) HandlerFunc(code string, fn HandlerFunc) {
	c.Lock()
	defer c.Unlock()

	c.path[code] = fn
}

func (c *Router) Execute(cmd string, args []byte, conn net.Conn) {
	c.RLock()
	defer c.RUnlock()

	if fn, ok := c.path[cmd]; ok {
		fn(conn, args)
	}
}

// Default parser
func ParseDefault(payload []byte) (string, []byte) {
	for i := 0; i < len(payload); i++ {
		if payload[i] == ' ' {
			return string(payload[:i]), payload[i+1:]
		}
	}

	return string(payload), []byte{}
}

func ParserJSON(payload []byte) (string, []byte) {
	for i := 0; i < len(payload); i++ {
		if payload[i] == 10 {
			return string(payload[:i]), payload[i+1:]
		}
	}

	return string(payload), []byte{}
}
