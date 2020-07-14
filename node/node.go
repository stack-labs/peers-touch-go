package node

import (
	"context"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	core "github.com/libp2p/go-libp2p-core"
	"sync"

	ipfsCore "github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
)

type Node interface {
	IPFS() coreapi.CoreAPI
	ID() PeerID
	Connect(peerInfo PeerInfo) error
	Touch() error
	Close()
}

type PeerInfo = core.PeerAddrInfo
type PeerID = core.PeerID

type node struct {
	ctx  context.Context
	ipfs coreapi.CoreAPI

	muPubSub sync.RWMutex
	muIPFS   sync.RWMutex
}

func (n *node) Connect(peerInfo PeerInfo) error {
	return n.ipfs.Swarm().Connect(n.ctx, peerInfo)
}

func (n *node) ID() string {
	panic("implement me")
}

func (n *node) Close() {
	panic("implement me")
}

func (n *node) IPFS() coreapi.CoreAPI {
	n.muIPFS.RLock()
	defer n.muIPFS.RUnlock()

	return n.ipfs
}

func NewNode(ctx context.Context, ipfs coreapi.CoreAPI, options ...Option) (n Node, err error) {
	opts := &Options{}
	for _, o := range options {
		o(opts)
	}

	n = newNode(ctx, ipfs, opts)

	return
}

func newNode(ctx context.Context, ipfs coreapi.CoreAPI, options *Options) *node {
	n := &node{
		ipfs: ipfs,
	}

	return n
}

type cleanFunc func()

func newIPFSNode(ctx context.Context, option libp2p.HostOption) (n *ipfsCore.IpfsNode, cf cleanFunc, err error) {
	core, err := ipfsCore.NewNode(ctx, &ipfsCore.BuildCfg{
		Online: true,
		Host:   option,
		ExtraOpts: map[string]bool{
			"pubsub": true,
		},
	})

	if err != nil {
		return
	}

	return core, func() {
		core.Close()
	}, nil
}
