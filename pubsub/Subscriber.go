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

	}
}

func NewSubscriber(ctx context.Context, ipfs coreapi.CoreAPI, topic string, opts ...SubOption) (sub Subscriber, err error) {
	options := &SubOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Topic == "" {
		return nil, fmt.Errorf("wrong empty topic")
	}

	pubSubSub, err := ipfs.PubSub().Subscribe(ctx, topic)
	if err != nil {
		return nil, err
	}

	id, err := ipfs.Key().Self(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get id for user")
	}

	s := &subscriber{
		ipfs:       coreapi.CoreAPI{},
		ipfsPubSub: pubSubSub,
		peerID:     id.ID(),
	}

	go s.start()

	sub = s
	return
}
