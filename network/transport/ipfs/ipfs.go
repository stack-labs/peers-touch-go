package ipfs

import (
	"fmt"
	"github.com/joincloud/peers-touch-go/codec"
	"time"

	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/network/transport"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

type ipfsTransport struct {
	opts       transport.Options
	host       host.Host
	protocolID protocol.ID
	codec      codec.Codec
}

func (i *ipfsTransport) Init(opts ...transport.Option) (err error) {
	for _, o := range opts {
		o(&i.opts)
	}

	defer func() {
		if err != nil {
			log.Errorf("transport init error: %s", i.String(), err)
		}
	}()

	pid, ok := i.opts.Context.Value(protocolIDKey{}).(protocol.ID)
	if !ok {
		err = fmt.Errorf("transport protocolID shouldn't be nil")
		return
	}

	i.protocolID = pid
	return nil
}

func (i *ipfsTransport) Options() transport.Options {
	return i.opts
}

func (i *ipfsTransport) Dial(addr string, opts ...transport.DialOption) (c transport.Client, err error) {
	options := transport.DialOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	defer func() {
		if err != nil {
			log.Errorf("transport %s dial peer %s error: %s", i.String(), addr, err)
		}
	}()

	if options.Timeout == 0 {
		options.Timeout = 5 * time.Second
	}

	// todo TLS
	id, err := peer.IDFromString(addr)
	if err != nil {
		err = fmt.Errorf("transport needs a legal addr error: %s", err)
		return
	}

	conn, err := i.host.NewStream(options.Context, id, i.protocolID)
	if err != nil {
		err = fmt.Errorf("transport %s dial peer %s error: %s", i.String(), addr, err)
		return
	}

	c = &ipfsTransportClient{
		&ipfsTransportSocket{
			stream: conn,
			local:  conn.Conn().LocalPeer().String(),
			remote: conn.Conn().RemotePeer().String(),
		},
	}

	return
}

func (i *ipfsTransport) Listen(addr string, opts ...transport.ListenOption) (tl transport.Listener, err error) {
	var options transport.ListenOptions
	for _, o := range opts {
		o(&options)
	}

	defer func() {
		if err != nil {
			log.Errorf("listen on addr: %s error: %s", addr, err)
		}
	}()

	h, err := libp2p.New(options.Context, libp2p.ListenAddrStrings(addr))
	if err != nil {
		err = fmt.Errorf("create new host on addr error: %s", addr)
		return nil, err
	}

	i.host = h

	return &ipfsTransportListener{
		host:  h,
		pid:   i.protocolID,
		opts:  options,
		codec: i.codec,
	}, nil
}

func (i *ipfsTransport) String() string {
	return "ipfs"
}
