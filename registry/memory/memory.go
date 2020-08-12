package memory

import (
	"fmt"
	"sync"

	"github.com/joincloud/peers-touch-go/registry"
)

type Registry struct {
	options registry.Options

	sync.RWMutex
	nodes map[string]*registry.Node
}

func (r *Registry) Register(node *registry.Node, opts ...registry.RegisterOption) error {
	r.Lock()
	defer r.Unlock()

	var options registry.RegisterOptions
	for _, o := range opts {
		o(&options)
	}

	if _, ok := r.nodes[node.Name]; ok {
		return fmt.Errorf("registry the node %s existed", node.Name)
	}

	r.nodes[node.Name] = node

	return nil
}

func (r *Registry) Deregister(node *registry.Node) error {
	if _, ok := r.nodes[node.Name]; ok {
		delete(r.nodes, node.Name)
	}

	return nil
}

func (r *Registry) GetNode(opts ...registry.NodeOption) (node *registry.Node, err error) {
	var options registry.NodeOptions
	for _, o := range opts {
		o(&options)
	}

	for _, node := range r.nodes {
		if options.Name == node.Name {
			return node, nil
		}
	}

	return nil, fmt.Errorf("there is no node: %s", options.Name)
}

func (r *Registry) ListNodes(opts ...registry.NodeOption) (nodes []*registry.Node, err error) {
	for _, node := range r.nodes {
		nodes = append(nodes, node)
	}

	return
}

func (r *Registry) Options() registry.Options {
	return r.options
}

func (r *Registry) Init(opts ...registry.Option) error {
	for _, o := range opts {
		o(&r.options)
	}

	r.Lock()
	defer r.Unlock()

	// todo load nodes from db or file
	r.nodes = map[string]*registry.Node{}

	return nil
}
