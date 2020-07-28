package registry

import (
	"context"
	"github.com/libp2p/go-libp2p-core/host"
	"time"
)

type Registry interface {
	Register(node *Node, opts ...RegisterOption) error
	Deregister(node *Node) error
	GetNode(...NodeOption) (node *Node, err error)
	ListNodes(...NodeOption) (nodes []*Node, err error)
	Init(...Option) error
	Options() Options
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type NodeOption func(*NodeOptions)

var (
	DefaultRegistry Registry
)

func Register(node *Node, opts ...RegisterOption) error {
	return DefaultRegistry.Register(node, opts...)
}

func Deregister(node *Node) error {
	return DefaultRegistry.Deregister(node)
}

func ListNodes(opts ...NodeOption) (nodes []*Node, err error) {
	return DefaultRegistry.ListNodes(opts...)
}

func RegisterContext(ctx context.Context) RegisterOption {
	return func(options *RegisterOptions) {
		options.Context = ctx
	}
}

func RegisterHost(h host.Host) RegisterOption {
	return func(options *RegisterOptions) {
		options.Host = h
	}
}

func RegisterInterval(duration time.Duration) RegisterOption {
	return func(options *RegisterOptions) {
		options.Interval = duration
	}
}

func RegisterTTL(ttl time.Duration) RegisterOption {
	return func(options *RegisterOptions) {
		options.TTL = ttl
	}
}
