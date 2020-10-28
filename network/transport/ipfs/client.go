package ipfs

import (
	"github.com/joincloud/peers-touch-go/network/transport"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
)

type ipfsTransportClient struct {
	host     host.Host
	protocol protocol.ID
	conn     network.Conn
	local    string
	remote   string
}

func (i *ipfsTransportClient) Recv(message *transport.Message) error {
	panic("implement me")
}

func (i *ipfsTransportClient) Send(message *transport.Message) error {
	panic("implement me")
}

func (i *ipfsTransportClient) Close() error {
	panic("implement me")
}

func (i *ipfsTransportClient) Local() string {
	return i.local
}

func (i *ipfsTransportClient) Remote() string {
	return i.remote
}
