package peer

import (
	core "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

type PeerAddrInfo = core.PeerAddrInfo
type PeerMultiAddr = ma.Multiaddr
type PeerID = core.PeerID
type Host = host.Host
type AddrInfo peer.AddrInfo
type Stream network.Stream

type Peer struct {
}

func NewPeer() *Peer {
	p := &Peer{}

	return p
}
