package node

import (
	"context"
	"sync"

	ipfsCore "github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/joincloud/peers-touch/peer"
	"github.com/joincloud/peers-touch/pubsub"
)

type Node interface {
	IPFS() coreapi.CoreAPI
	ID() peer.PeerID
	Broker() pubsub.Broker
	Connect(peerInfo peer.PeerInfo) error
	Disconnect(multiAddr peer.PeerMultiAddr) error
	// Touch the file#subjectID from peer#peerID
	Touch(peerID peer.PeerID, subjectID string) ([]byte, error)
	Close()
}

type node struct {
	ctx      context.Context
	ipfs     coreapi.CoreAPI
	id       peer.PeerID
	broker   pubsub.Broker
	muPubSub sync.RWMutex
	muIPFS   sync.RWMutex
}

func (n *node) Broker() pubsub.Broker {
	return n.broker
}

func (n *node) Touch() error {
	panic("implement me")
}

func (n *node) Connect(peerInfo peer.PeerInfo) error {
	return n.ipfs.Swarm().Connect(n.ctx, peerInfo)
}

func (n *node) Disconnect(multiAddr peer.PeerMultiAddr) error {
	return n.ipfs.Swarm().Disconnect(n.ctx, multiAddr)
}

func (n *node) ID() peer.PeerID {
	return n.id
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
