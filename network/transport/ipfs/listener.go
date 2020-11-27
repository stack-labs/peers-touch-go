package ipfs

import (
	"bufio"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/stack-labs/peers-touch-go/codec"
	"github.com/stack-labs/peers-touch-go/network/transport"
)

type ipfsTransportListener struct {
	host  host.Host
	pid   protocol.ID
	opts  transport.ListenOptions
	codec codec.Codec
}

func (i *ipfsTransportListener) Addr() string {
	return i.host.Addrs()[0].String()
}

func (i *ipfsTransportListener) Close() error {
	i.host.RemoveStreamHandler(i.pid)
	return nil
}

func (i *ipfsTransportListener) Accept(fn func(socket transport.Socket)) error {
	i.host.SetStreamHandler(i.pid, func(s network.Stream) {
		fn(&ipfsTransportSocket{
			stream: s,
			r:      bufio.NewReader(s),
			w:      bufio.NewWriter(s),
			codec:  i.codec,
		})
	})
	return nil
}
