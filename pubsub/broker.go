package pubsub

import (
	"context"
	"github.com/joincloud/peers-touch-go/logger"
	"sync"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch-go/codec"
	"github.com/pkg/errors"
)

type Broker interface {
	Init(...BrokerOption) error
	Pub(ctx context.Context, event Event) error
	Sub(ctx context.Context, topic string, handler Handler) (Subscriber, error)
	Codec() codec.Codec
	Unsub(topic string) error
	Close() error
}

type Message struct {
	Header map[string]string
	Body   interface{}
}

type Event interface {
	Name() string
	Message() Message
}

type Handler func(topic Event)

type event struct {
	N string
	M Message
}

func (e *event) Name() string {
	return e.N
}

func (e *event) Message() Message {
	return e.M
}

type broker struct {
	codec       codec.Codec
	coreAPI     iface.CoreAPI
	subscribers map[string]Subscriber
	muMux       sync.RWMutex
	exit        chan chan error
}

func (b *broker) Init(...BrokerOption) error {
	return nil
}

func (b *broker) Codec() codec.Codec {
	return b.codec
}

func (b *broker) Pub(ctx context.Context, event Event) (err error) {
	bytes, err := b.codec.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "unable to marshal message")
	}
	id, err := b.coreAPI.Key().Self(ctx)
	if err != nil {
		panic(err)
	}

	logger.Debugf("pub id: %s", id.ID())
	err = b.coreAPI.PubSub().Publish(ctx, event.Name(), bytes)
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

	if bo.codec == nil {
		bo.codec = codec.Codecs["json"]()
	}

	b := &broker{
		codec:       bo.codec,
		coreAPI:     bo.coreAPI,
		subscribers: map[string]Subscriber{},
		exit:        make(chan chan error),
	}

	return b
}

func NewEvent(name string, m Message) Event {
	return &event{
		N: name,
		M: m,
	}
}
