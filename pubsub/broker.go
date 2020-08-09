package pubsub

import (
	"context"
	"sync"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch-go/codec"
	"github.com/pkg/errors"
)

type Broker interface {
	Pub(ctx context.Context, event Topic) error
	Sub(ctx context.Context, topic string, handler Handler) (Subscriber, error)
	Unsub(topic string) error
	Close() error
}

type Message struct {
	Header map[string]string
	Body   []byte
}

type Topic interface {
	Name() string
	Message() Message
	Bytes() ([]byte, error)
}

type Handler func(topic Topic)

type topic struct {
	n string
	m Message
	c codec.Codec
}

func (t *topic) Name() string {
	return t.n
}

func (t *topic) Message() Message {
	return t.m
}

func (t *topic) Bytes() ([]byte, error) {
	return t.c.Marshal(t.m)
}

type broker struct {
	coreAPI     iface.CoreAPI
	subscribers map[string]Subscriber
	muMux       sync.RWMutex
	exit        chan chan error
}

func (b *broker) Pub(ctx context.Context, topic Topic) (err error) {
	bytes, err := topic.Bytes()
	if err != nil {
		return errors.Wrap(err, "unable to marshal message")
	}

	err = b.coreAPI.PubSub().Publish(ctx, topic.Name(), bytes)
	if err != nil {
		return errors.Wrap(err, "unable to publish data on pubsub")
	}

	return
}

func (b *broker) Sub(ctx context.Context, topic string, handler Handler) (Subscriber, error) {
	b.muMux.Lock()
	defer b.muMux.Unlock()

	if sub, ok := b.subscribers[topic]; ok {
		return sub, nil
	}

	s, _ := NewSubscriber(ctx,
		SubTopic(topic),
		SubCoreAPI(b.coreAPI),
		SubHandler(handler),
	)

	b.subscribers[topic] = s

	return s, nil
}

func (b *broker) Unsub(topic string) error {
	b.muMux.Lock()
	defer b.muMux.Unlock()

	delete(b.subscribers, topic)

	return nil
}

func (b *broker) Close() error {
	ch := make(chan error)
	b.exit <- ch
	err := <-ch
	return err
}

func NewBroker(options ...BrokerOption) Broker {
	bo := &BrokerOptions{}
	for _, option := range options {
		option(bo)
	}

	b := &broker{
		coreAPI:     bo.coreAPI,
		subscribers: map[string]Subscriber{},
		exit:        make(chan chan error),
	}

	return b
}

func NewTopic(name string, m Message, codec codec.Codec) Topic {
	return &topic{
		n: name,
		m: m,
		c: codec,
	}
}
