package pubsub

import (
	"context"
	"fmt"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/joincloud/peers-touch/peer"
	"github.com/pkg/errors"
)

type Subscriber interface {
	Topic() string
	Unsubscribe() error
}

type subscriber struct {
	ipfs   coreapi.CoreAPI
	peerID peer.PeerID
}

func (s *subscriber) Topic() string {
	panic("implement me")
}

func (s *subscriber) Unsubscribe() error {
	panic("implement me")
}

func (s *subscriber) start() {

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

	sub = &subscriber{
		ipfs:   coreapi.CoreAPI{},
		peerID: "",
	}

	return
}
