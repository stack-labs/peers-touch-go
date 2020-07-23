package pubsub

import (
	"github.com/ipfs/go-ipfs/core/coreapi"
	iface "github.com/ipfs/interface-go-ipfs-core"
)

type BrokerOptions struct {
	coreAPI iface.CoreAPI
}

type BrokerOption func(o *BrokerOptions)

func BrokerCoreAPI(coreAPI iface.CoreAPI) BrokerOption {
	return func(o *BrokerOptions) {
		o.coreAPI = coreAPI
	}
}

type SubOptions struct {
	Topic   string
	coreAPI *coreapi.CoreAPI
	Handler Handler
}

func NewSubOptions() SubOptions {
	return SubOptions{}
}

type SubOption func(o *SubOptions)

func SubCoreAPI(coreAPI *coreapi.CoreAPI) SubOption {
	return func(o *SubOptions) {
		o.coreAPI = coreAPI
	}
}

func SubTopic(topic string) SubOption {
	return func(o *SubOptions) {
		o.Topic = topic
	}
}

func SubHandler(handler Handler) SubOption {
	return func(o *SubOptions) {
		o.Handler = handler
	}
}
