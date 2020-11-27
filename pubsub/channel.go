package pubsub

import (
	"context"
	"github.com/stack-labs/peers-touch-go/codec"

	"github.com/libp2p/go-libp2p-pubsub"
	log "github.com/stack-labs/peers-touch-go/logger"
	"github.com/stack-labs/peers-touch-go/peer"
)

var (
	DefaultChatMsgBufSize = 128
)

type Channel struct {
	Messages chan *Message
	name     string
	ctx      context.Context
	ps       *pubsub.PubSub
	topic    *pubsub.Topic
	sub      *pubsub.Subscription
	peerID   peer.PeerID
	codec    codec.Codec
	metadata map[string]string
}

func (ch *Channel) Ctx() context.Context {
	return ch.ctx
}

func (ch *Channel) Name() string {
	return ch.name
}

func (ch *Channel) Metadata() map[string]string {
	return ch.metadata
}

func joinChannel(ctx context.Context, opts ...ChannelOption) (ch *Channel, err error) {
	defer func() {
		if err != nil {
			log.Errorf("join channel error: %s", err)
		}
	}()

	options := &ChannelOptions{}
	for _, opt := range opts {
		opt(options)
	}

	// todo check options
	log.Infof("join channel: %s with name: %v", options.Name, options.Metadata)
	topic, err := options.Pubsub.Join(options.Name)
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	ch = &Channel{
		ctx:    ctx,
		ps:     options.Pubsub,
		topic:  topic,
		sub:    sub,
		peerID: options.PeerID,
		name:   options.Name,
		codec:  options.Codec,
		// todo bufSize option
		Messages: make(chan *Message, DefaultChatMsgBufSize),
	}

	go ch.listen()
	return ch, nil
}

func (ch *Channel) listen() {
	for {
		msg, err := ch.sub.Next(ch.ctx)
		if err != nil {
			close(ch.Messages)
			return
		}
		// todo self msg manage
		// ignore the msg from self node
		if msg.ReceivedFrom == ch.peerID {
			log.Debugf("ignore the msg from self node: %s", ch.peerID)
			continue
		}
		cm := &Message{}
		err = ch.codec.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}

		ch.Messages <- cm
	}
}

func (ch *Channel) ListPeers() []peer.PeerID {
	return ch.ps.ListPeers(ch.name)
}

func (ch *Channel) Publish(msg string) error {
	m := Message{
		Body: msg,
		// todo metadata should not be transfered directly
		Header: ch.metadata,
	}

	bytes, err := ch.codec.Marshal(m)
	if err != nil {
		return err
	}
	return ch.topic.Publish(ch.ctx, bytes)
}
