package pubsub

import (
	"context"
	"encoding/json"

	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/peer"
	"github.com/libp2p/go-libp2p-pubsub"
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

func JoinChannel(ctx context.Context, opts ...ChannelOption) (ch *Channel, err error) {
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
		// todo use codec
		err = json.Unmarshal(msg.Data, cm)
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
		Header: map[string]string{
			// todo, more header option
			"content-type": "application/txt",
			"from":         ch.peerID.Pretty(),
		},
	}

	// todo, use codec option
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ch.topic.Publish(ch.ctx, bytes)
}
