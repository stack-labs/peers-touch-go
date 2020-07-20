package pubsub

import "github.com/ipfs/go-ipfs/core/coreapi"

type BrokerOptions struct {
	coreAPI *coreapi.CoreAPI
}

type BrokerOption func(o *BrokerOptions)

func BrokerCoreAPI(coreAPI *coreapi.CoreAPI) BrokerOption {
	return func(o *BrokerOptions) {
		o.coreAPI = coreAPI
	}
}

type SubOptions struct {
	Topic   string
	coreAPI *coreapi.CoreAPI
}

type SubOption func(o *SubOptions)

func SubCoreAPI(coreAPI *coreapi.CoreAPI) SubOption {
	return func(o *SubOptions) {
		o.coreAPI = coreAPI
	}
}
