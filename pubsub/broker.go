package pubsub

import (
	"context"
	"fmt"
	"sync"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch-go/codec"
	"github.com/joincloud/peers-touch-go/logger"
	"github.com/pkg/errors"
)

type Broker interface {
	Init(...BrokerOption) error
	Pub(ctx context.Context, event Event) error
	Sub(ctx context.Context, topic string, handler Handler) (Subscriber, error)
	Unsub(topic string) error
	Join(ctx context.Context, opts ...ChannelOption) (ch *Channel, err error)
	// Connect(id peer.PeerID) (err error)
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
	codec       codec.Codec
	coreAPI     iface.CoreAPI
	subscribers map[string]Subscriber
	muMux       sync.RWMutex
	exit        chan chan error
	chans       map[string]*Channel
}

func (b *broker) Join(ctx context.Context, opts ...ChannelOption) (ch *Channel, err error) {
	ch, err = joinChannel(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("join channel err: %s", err)
	}
	b.chans[ch.name] = ch
	return
}

func (b *broker) Init(...BrokerOption) error {
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

	logger.Debugf("pub id: %s", id.ID())
	err = b.coreAPI.PubSub().Publish(ctx, event.Name(), bytes)
	if err != nil {
		return errors.Wrap(err, "unable to publish data on pubsub")
	}

	return
}

/*func (b *broker) Connect(id peer.PeerID) (err error) {
	// Turn the destination into a multiaddr.
	maddr, err := multiaddr.NewMultiaddr(*dest)
	if err != nil {
		log.Fatalln(err)
	}

	// Extract the peer ID from the multiaddr.
	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		log.Fatalln(err)
	}

	// Add the destination's peer multiaddress in the peerstore.
	// This will be used during connection and stream creation by libp2p.
	host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	s, err := host.NewStream(context.Background(), info.ID, "/chat/1.0.0")
	if err != nil {
		panic(err)
	}

	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
}*/

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
