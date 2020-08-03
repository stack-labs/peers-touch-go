package node

import (
	"context"
	"sync"

	ipfsCore "github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch-go/file"
	"github.com/joincloud/peers-touch-go/peer"
	"github.com/joincloud/peers-touch-go/pubsub"
	golib "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	peer2 "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
)

var (
	DefaultAddrs = []string{"/ip4/0.0.0.0/tcp/0"}
)

type Node interface {
	IPFS() iface.CoreAPI
	ID() peer.PeerID
	Broker() pubsub.Broker
	Connect(peerInfo peer.PeerAddrInfo) error
	Disconnect(multiAddr peer.PeerMultiAddr) error
	// Touch the file#subjectID from peer#peerID
	Touch(peerID peer.PeerID, subjectID string) (file.File, error)
	Close()
}

type node struct {
	ctx      context.Context
	ipfs     iface.CoreAPI
	id       peer.PeerID
	broker   pubsub.Broker
	host     peer.Host
	muPubSub sync.RWMutex
	muIPFS   sync.RWMutex
}

func (n *node) Touch(peerID peer.PeerID, subjectID string) (file.File, error) {
	panic("implement me")
}

func (n *node) Broker() pubsub.Broker {
	return n.broker
}

func (n *node) Connect(peerInfo peer.PeerAddrInfo) error {
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

func (n *node) IPFS() iface.CoreAPI {
	n.muIPFS.RLock()
	defer n.muIPFS.RUnlock()

	return n.ipfs
}

func NewNode(ctx context.Context, options ...Option) (n Node, err error) {
	opts := &Options{}
	for _, o := range options {
		o(opts)
	}

	if opts.Adds == nil {
		opts.Adds = DefaultAddrs
	}

	if opts.Host == nil {
		opts.Host, err = golib.New(ctx, golib.ListenAddrStrings(opts.Adds...))
		if err != nil {
			panic(err)
		}
	}

	if opts.IPFS == nil {
		n, _, err := newIPFSNode(ctx, func(ctx context.Context, id peer2.ID, ps peerstore.Peerstore, options ...golib.Option) (host host.Host, err error) {
			return opts.Host, nil
		})
		if opts.IPFS, err = coreapi.NewCoreAPI(n); err != nil {
			panic(err)
		}
	}

	if opts.Broker == nil {
		opts.Broker = pubsub.NewBroker(pubsub.BrokerCoreAPI(opts.IPFS))
	}

	n = newNode(ctx, opts)

	return
}

func newNode(ctx context.Context, options *Options) *node {
	n := &node{
		ipfs:   options.IPFS,
		broker: options.Broker,
		id:     options.PeerID,
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
