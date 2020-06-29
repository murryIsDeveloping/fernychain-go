package p2p

// P2P represents the p2p server
type P2P struct {
	peers []string
}

func (p2p *P2P) addPeer(peer string) {
	p2p.peers = append(p2p.peers, peer)
}

func (p2p *P2P) removePeer(peer string) {
	for i := 0; i < len(p2p.peers); i++ {
		if p2p.peers[i] == peer {
			p2p.peers[i] = p2p.peers[len(p2p.peers)-1]
			p2p.peers[len(p2p.peers)-1] = ""
			p2p.peers = p2p.peers[:len(p2p.peers)-1]
		}
	}
}
