package ipfs

import "github.com/joincloud/peers-touch-go/network/transport"

type ipfsTransportListener struct {
	it *ipfsTransport
}

func (i *ipfsTransportListener) Addr() string {
	panic("implement me")
}

func (i *ipfsTransportListener) Close() error {
	panic("implement me")
}

func (i *ipfsTransportListener) Accept(f func(socket transport.Socket)) error {
	panic("implement me")
}
