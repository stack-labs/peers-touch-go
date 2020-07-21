package pubsub

import (
	"context"
	"github.com/pkg/errors"
	"sync"

	"github.com/ipfs/go-ipfs/core/coreapi"
)

type Broker interface {
	Pub(ctx context.Context, event Event) error
	Sub(ctx context.Context, topic string) (Subscriber, error)
	Unsub(topic string) error
	Close() error
}

type Message struct {
	Header map[string]string
	Body   []byte
}

func (m *Message) Bytes() []byte {
	return nil
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

func (m *event) Topic() string {
	panic("implement me")
}

func (m *event) Message() Message {
	panic("implement me")
}

type broker struct {
	coreAPI     *coreapi.CoreAPI
	subscribers map[string]Subscriber
	muMux       sync.RWMutex
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

	s, _ := NewSubscription(ctx, b.coreAPI, topic, &SubOptions{
		Topic:   topic,
		coreAPI: b.coreAPI,
		Handler: handler,
	})

	b.subscribers[topic] = s

	return s, nil
}

func (b *broker) Unsub(topic string) error {
	panic("implement me")
}

func (b *broker) Close() error {
	panic("implement me")
}

func NewBroker(options ...BrokerOption) Broker {
	bo := &BrokerOptions{}
	for _, option := range options {
		option(bo)
	}

	b := &broker{
		coreAPI: bo.coreAPI,
	}

	return b
}

func NewSubscription(ctx context.Context, api *coreapi.CoreAPI, topic string, options *SubOptions) (s Subscriber, err error) {

}

func NewEvent(topic string, content []byte) Event {
	return &event{
		t: topic,
		m: &Message{
			Header: nil,
			Body:   content,
		}
	}
}
