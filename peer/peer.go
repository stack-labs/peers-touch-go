package peer

import (
	core "github.com/libp2p/go-libp2p-core"
	ma "github.com/multiformats/go-multiaddr"
)

type PeerInfo = core.PeerAddrInfo
type PeerMultiAddr = ma.Multiaddr
type PeerID = core.PeerID
