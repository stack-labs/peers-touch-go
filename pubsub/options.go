package pubsub

import (
	"context"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch-go/codec"
	"github.com/joincloud/peers-touch-go/peer"
	"github.com/libp2p/go-libp2p-pubsub"
)

type BrokerOptions struct {
	coreAPI iface.CoreAPI
	codec   codec.Codec
}

type BrokerOption func(o *BrokerOptions)

func BrokerCoreAPI(coreAPI iface.CoreAPI) BrokerOption {
	return func(o *BrokerOptions) {
		o.coreAPI = coreAPI
	}
}

func BrokerCodec(codec codec.Codec) BrokerOption {
	return func(o *BrokerOptions) {
		o.codec = codec
	}
}

type SubOptions struct {
	Topic   string
	codec   codec.Codec
	coreAPI iface.CoreAPI
	Handler Handler
}

func NewSubOptions() SubOptions {
	return SubOptions{}
}

type SubOption func(o *SubOptions)

func SubCoreAPI(coreAPI iface.CoreAPI) SubOption {
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

func SubCodec(codec codec.Codec) SubOption {
	return func(o *SubOptions) {
		o.codec = codec
	}
}

type ChannelOptions struct {
	Name   string
	Topic  string
	Ctx    context.Context
	Pubsub *pubsub.PubSub
	PeerID peer.PeerID
	// todo metadata struct
	Metadata map[string]string
	Codec    codec.Codec
}

type ChannelOption func(o *ChannelOptions)

func ChannelName(name string) ChannelOption {
	return func(o *ChannelOptions) {
		o.Name = name
	}
}

func ChannelCtx(ctx context.Context) ChannelOption {
	return func(o *ChannelOptions) {
		o.Ctx = ctx
	}
}

func ChannelTopic(topic string) ChannelOption {
	return func(o *ChannelOptions) {
		o.Topic = topic
	}
}

func ChannelPeerID(peerID peer.PeerID) ChannelOption {
	return func(o *ChannelOptions) {
		o.PeerID = peerID
	}
}

func ChannelMetadata(metadata map[string]string) ChannelOption {
	return func(o *ChannelOptions) {
		o.Metadata = metadata
	}
}

func ChannelPubsub(pubsub *pubsub.PubSub) ChannelOption {
	return func(o *ChannelOptions) {
		o.Pubsub = pubsub
	}
}

func ChannelCodec(cc codec.Codec) ChannelOption {
	return func(o *ChannelOptions) {
		o.Codec = cc
	}
}

type PushOptions struct {
	// todo options
}

type PushOption func(o *PushOptions)
