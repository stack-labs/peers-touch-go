package mdns

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/joincloud/peers-touch－go/registry"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "pubsub-chat-example"

func init() {
	registry.Registries["mdns"] = NewRegistry
}

type mdns struct {
	discoverMap map[string]discovery.Service

	mtx sync.RWMutex
}

func (m *mdns) String() string {
	return "mdns"
}

func (m *mdns) Register(node *registry.Node, opts ...registry.RegisterOption) (err error) {
	options := &registry.RegisterOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Context == nil {
		options.Context = context.Background()
	}

	if options.Host == nil {
		return fmt.Errorf("Host is nil ")
	}

	if options.Interval == 0 {
		options.Interval = DiscoveryInterval
	}

	if options.ServiceTag == "" {
		options.ServiceTag = DiscoveryServiceTag
	}

	disc, err := discovery.NewMdnsService(options.Context, options.Host, options.Interval, options.ServiceTag)
	if err != nil {
		return fmt.Errorf("discover register error: %s", err)
	}
	if _, ok := m.discoverMap[node.Name]; !ok {
		m.discoverMap[node.Name] = disc
	}

	node.Host = options.Host
	disc.RegisterNotifee(node)

	return
}

func (m *mdns) Deregister(node *registry.Node) error {
	// todo
	if disc, ok := m.discoverMap[node.Name]; ok {
		disc.UnregisterNotifee(node)
	}

	return nil
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

func NewRegistry(options ...registry.Option) registry.Registry {
	m := &mdns{
		discoverMap: make(map[string]discovery.Service),
	}

	return m
}
