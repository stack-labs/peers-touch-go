package pubsub

import (
	"context"
	"fmt"

	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch/peer"
	"github.com/pkg/errors"
)

type Subscriber interface {
	Topic() string
	Unsubscribe() error
}

type subscriber struct {
	opts       *SubOptions
	ipfs       coreapi.CoreAPI
	ipfsPubSub iface.PubSubSubscription
	peerID     peer.PeerID
	handler    Handler
}

func (s *subscriber) Topic() string {
	return s.opts.Topic
}

func (s *subscriber) Unsubscribe() error {
	return s.ipfsPubSub.Close()
}

func (s *subscriber) start(ctx context.Context) {
	for {
		msg, err := s.ipfsPubSub.Next(ctx)
		if err != nil {
			// todo error
		}

		// ignore self msg
		if msg.From() == s.peerID {
			continue
		}

		topic := msg.Topics()[0]
		if topic != s.opts.Topic {
			continue
		}

		evt := NewEvent(s.Topic(), msg.Data())
		s.handler(evt)
	}
}

func NewSubscriber(ctx context.Context, opts ...SubOption) (sub Subscriber, err error) {
	options := &SubOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Topic == "" {
		return nil, fmt.Errorf("wrong empty topic")
	}

	pubSubSub, err := options.coreAPI.PubSub().Subscribe(ctx, options.Topic)
	if err != nil {
		return nil, err
	}

	id, err := options.coreAPI.Key().Self(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get id for user")
	}

	s := &subscriber{
		ipfs:       coreapi.CoreAPI{},
		ipfsPubSub: pubSubSub,
		peerID:     id.ID(),
		handler:    options.Handler,
	}

	go s.start(ctx)

	sub = s
	return
}
