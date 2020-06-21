package platform

import "github.com/joincloud/peers-touch/registry"

type Platform interface {
	Init(opts ...InitOption) error
	Connect(opts ...ConnectOption) error
	Close() error
}

type InitOptions struct {
	URL  string
	Node registry.Node
}

type InitOption func(*InitOptions)

type ConnectOptions struct {
	URL  string
	Node registry.Node
}

type ConnectOption func(*InitOptions)
