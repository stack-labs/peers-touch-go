package ipfs

import (
	"context"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/stack-labs/peers-touch-go/network/transport"
)

type protocolIDKey struct{}

func ProtocolID(pid protocol.ID) transport.Option {
	return func(o *transport.Options) {
		o.Context = context.WithValue(o.Context, protocolIDKey{}, pid)
	}
}
