package pubsub

import "github.com/ipfs/go-ipfs/core/coreapi"

type BrokerOptions struct {
	coreAPI coreapi.CoreAPI
}

type BrokerOption func(o *BrokerOptions)

func BrokerCoreAPI(coreAPI coreapi.CoreAPI) BrokerOption {
	return func(o *BrokerOptions) {
		o.coreAPI = coreAPI
	}
}
