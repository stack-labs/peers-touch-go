package pubsub

import (
	"context"
	"encoding/json"
	"sync"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/pkg/errors"
)

type Broker interface {
	Pub(ctx context.Context, event Event) error
	Sub(ctx context.Context, topic string, handler Handler) (Subscriber, error)
	Unsub(topic string) error
	Close() error
}

type Message struct {
	Header map[string]string
	Body   []byte
}

func (m Message) Bytes() []byte {
	// todo use codec
	msg, _ := json.Marshal(m)
	return msg
}

type Event interface {
	Topic() string
	Message() Message
}

type Handler func(event Event)

type event struct {
	t string
	m Message
}

func (e *event) Topic() string {
	return e.t
}

func (e *event) Message() Message {
	return e.m
}

type broker struct {
	coreAPI     iface.CoreAPI
	subscribers map[string]Subscriber
	muMux       sync.RWMutex
	exit        chan chan error
}

func (b *broker) Pub(ctx context.Context, event Event) (err error) {
	err = b.coreAPI.PubSub().Publish(ctx, event.Topic(), event.Message().Bytes())
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

func NewEvent(topic string, content []byte) Event {
	return &event{
		t: topic,
		m: Message{
			Header: nil,
			Body:   content,
		},
	}
}
