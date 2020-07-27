package node

import (
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch/peer"
)

type Options struct {
	Directory string
	Host      peer.Host
	ID        string
	IPFS      iface.CoreAPI
	PeerID    peer.PeerID
}

type Option func(options *Options)

func Host(host peer.Host) Option {
	return func(options *Options) {
		options.Host = host
	}
}

func ID(id string) Option {
	return func(options *Options) {
		options.ID = id
	}
}

func IPFS(is iface.CoreAPI) Option {
	return func(options *Options) {
		options.IPFS = is
	}
}
