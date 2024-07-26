package router

import "root/p2p"

func RouterP2P(p2p p2p.Peer) map[string]HandlerFunc {
	return map[string]HandlerFunc{
		"CanI":           p2p.CanI,
		"AddNode":        p2p.AddNode,
		"Info":           p2p.Info,
		"AllowToNetwork": p2p.AllowToNetwork,
		"ChangeProcess":  p2p.ChangeProcess,
	}
}
