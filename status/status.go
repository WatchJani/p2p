package status

import (
	"errors"
	"net"
	"sync"
)

type Status struct {
	InProcess bool
	Waiting   []net.Conn
	Current   string

	Approve map[net.Conn]struct{}
	sync.RWMutex
}

func (s *Status) AddToWaiting(conn net.Conn) {
	s.Lock()
	defer s.Unlock()

	s.Waiting = append(s.Waiting, conn)
}

func (s *Status) ChangeProgress() {
	s.Lock()
	defer s.Unlock()

	s.InProcess = !s.InProcess
}

func (s *Status) ApproveReset() {
	s.Approve = make(map[net.Conn]struct{})
}

func (s *Status) AddToApprove(conn net.Conn) error {
	if _, ok := s.Approve[conn]; ok {
		return errors.New("this node already send approvement")
	}

	s.Approve[conn] = struct{}{}
	return nil
}

func (s *Status) NumberOfApprovement() int {
	return len(s.Approve)
}

func (s Status) IsInProcess() bool {
	return s.InProcess
}

func New() Status {
	return Status{
		Waiting: make([]net.Conn, 0),
		Approve: make(map[net.Conn]struct{}),
	}
}
