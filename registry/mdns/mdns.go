package mdns

import (
	"github.com/joincloud/peers-touch－go/registry"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"time"
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "pubsub-chat-example"

type mdns struct {
	service discovery.Service
}

func (m *mdns) Register(node *registry.Node, opts ...registry.RegisterOption) error {

	panic("implement me")
}

func (m *mdns) Deregister(node *registry.Node) error {
	panic("implement me")
}

func (m *mdns) GetNode(...registry.NodeOption) (node *registry.Node, err error) {
	panic("implement me")
}

func (m *mdns) ListNodes(...registry.NodeOption) (nodes []*registry.Node, err error) {
	panic("implement me")
}

func (m *mdns) Init(...registry.Option) error {
	panic("implement me")
}

func (m *mdns) Options() registry.Options {
	panic("implement me")
}
