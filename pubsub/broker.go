package pubsub

import (
	"bufio"
	"context"
	"fmt"
	"sync"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/network"
	peerlib "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/stack-labs/peers-touch-go/codec"
	log "github.com/stack-labs/peers-touch-go/logger"
	"github.com/stack-labs/peers-touch-go/peer"
)

type Broker interface {
	Init(...BrokerOption) error
	Pub(ctx context.Context, event Event) error
	Sub(ctx context.Context, topic string, handler Handler) (Subscriber, error)
	Unsub(topic string) error
	Join(ctx context.Context, opts ...ChannelOption) (ch *Channel, err error)
	// Connect(id peer.PeerID) (err error)
	Touch(ctx context.Context, opts ...TouchOption) (err error)
	Codec() codec.Codec
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
	options     BrokerOptions
	codec       codec.Codec
	coreAPI     iface.CoreAPI
	host        peer.Host
	subscribers map[string]Subscriber
	muMux       sync.RWMutex
	exit        chan chan error
	chans       map[string]*Channel
}

func (b *broker) Touch(ctx context.Context, opts ...TouchOption) (err error) {
	options := TouchOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	defer func() {
		if err != nil {
			log.Errorf("touch dest error: %s", err)
		}
	}()

	// todo check more options
	if options.DestAddr == "" {
		b.host.SetStreamHandler("/chat/1.0.0", func(stream network.Stream) {
			options.StreamHandler(stream)
		}, /*  func(s network.Stream) {
			rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
			if options.Writer != nil {
				go options.Writer(rw)
			}
			if options.Reader != nil {
				go options.Reader(rw)
			}
		}*/)
	} else {
		maddr, err := multiaddr.NewMultiaddr(options.DestAddr)
		if err != nil {
			return fmt.Errorf("touch dest node err: %s", err)
		}

		info, err := peerlib.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			return fmt.Errorf("touch get dest node addr info err: %s", err)
		}

		ctx := context.WithValue(context.Background(), "test", "123")
		b.host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)
		s, err := b.host.NewStream(ctx, info.ID, "/chat/1.0.0")
		if err != nil {
			return fmt.Errorf("touch make new stream err: %s", err)
		}

		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		if options.Writer != nil {
			go options.Writer(rw)
		}

		if options.Reader != nil {
			go options.Reader(rw)
		}
	}

	return nil
}

func (b *broker) Join(ctx context.Context, opts ...ChannelOption) (ch *Channel, err error) {
	ch, err = joinChannel(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("join channel err: %s", err)
	}
	b.chans[ch.name] = ch
	return
}

func (b *broker) Init(opts ...BrokerOption) (err error) {
	defer func() {
		if err != nil {
			log.Errorf("init broker error: %s", err)
		}
	}()

	if b.options.host == nil {
		return fmt.Errorf("broker's host shouldn't be nil")
	}

	b.host = b.options.host
	b.chans = make(map[string]*Channel)
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

	log.Debugf("pub id: %s", id.ID())
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
	bo := BrokerOptions{}
	for _, option := range options {
		option(&bo)
	}

	if bo.codec == nil {
		bo.codec = codec.Codecs["json"]()
	}

	b := &broker{
		options:     bo,
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
