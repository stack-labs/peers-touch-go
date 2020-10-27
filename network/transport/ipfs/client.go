package ipfs

import (
	"github.com/joincloud/peers-touch-go/network/transport"
	"github.com/libp2p/go-libp2p-core/network"
)

type ipfsClient struct {
	conn network.Conn
}

func (i *ipfsClient) Recv(message *transport.Message) error {
	panic("implement me")
}

func (i *ipfsClient) Send(message *transport.Message) error {
	panic("implement me")
}

func (i *ipfsClient) Close() error {
	panic("implement me")
}

func (i *ipfsClient) Local() string {
	panic("implement me")
}

func (i *ipfsClient) Remote() string {
	panic("implement me")
}
