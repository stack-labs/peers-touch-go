package node

import (
	p2pcore "github.com/libp2p/go-libp2p-core"
)

type Options struct {
	ID        string
	PeerID    p2pcore.PeerID
	Directory string
}

type Option func(options *Options)
