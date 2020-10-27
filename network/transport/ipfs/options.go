package ipfs

import (
	"context"
	"github.com/joincloud/peers-touch-go/network/transport"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
)

type hostKey struct{}
type protocolIDKey struct{}

func Host(host host.Host) transport.Option {
	return func(o *transport.Options) {
		o.Context = context.WithValue(o.Context, hostKey{}, host)
	}
}

func ProtocolID(pid protocol.ID) transport.Option {
	return func(o *transport.Options) {
		o.Context = context.WithValue(o.Context, protocolIDKey{}, pid)
	}
}
