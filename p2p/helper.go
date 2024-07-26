package p2p

import "net"

func (p *Peer) Broadcast(msg string) error {
	for _, conn := range p.peer {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteMessage(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
}

func NewConn(addr string) (net.Conn, error) {
	return net.Dial("tcp", addr)
}
