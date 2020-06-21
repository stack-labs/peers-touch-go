package registry

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

func ListNodes(opts ...NodeOption) (nodes []*Node, err error) {
	return DefaultRegistry.ListNodes(opts...)
}
