package ipfs

import (
	"github.com/joincloud/peers-touch-go/network/transport"
)

type ipfsTransportListener struct {
	it   *ipfsTransport
	opts transport.ListenOptions
}

func (i *ipfsTransportListener) Addr() string {
	return i.it.host.Addrs()[0].String()
}

func (i *ipfsTransportListener) Close() error {
	return i.it.host.Close()
}

func (i *ipfsTransportListener) Accept(f func(socket transport.Socket)) error {
	for {

	}
}
