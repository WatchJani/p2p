package server

import (
	"fmt"
	"log"
	"net"
	"root/p2p"
	"root/router"
)

type Server struct {
	conn chan net.Conn
	addr string
	p2p.Peer
	router.Router
}

func New(addr string) *Server {
	connCh := make(chan net.Conn)
	p2p := p2p.New(connCh)

	s := &Server{
		addr:   addr,
		conn:   connCh,
		Peer:   p2p,
		Router: router.New(p2p),
	}

	go s.ConnectionReceive()

	return s
}

func (s *Server) Listen() error {
	ls, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		s.conn <- conn
	}
}

func (s *Server) Handler(conn net.Conn) {
	buffer := make([]byte, 4096)
	defer conn.Close()

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			break
		}

		cmd, args := s.Parse(buffer[:n])
		s.Execute(cmd, args, conn)
	}

	fmt.Printf("Connection is close | ip: %s\n", conn.RemoteAddr().String())
}

// client and server connection
func (s *Server) ConnectionReceive() {
	for {
		conn := <-s.conn
		go s.Handler(conn)
	}
}
