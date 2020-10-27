package ipfs

import (
	"fmt"
	"time"

	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/network/transport"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
)

type ipfsTransport struct {
	opts       transport.Options
	host       host.Host
	protocolID protocol.ID
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

	h, ok := i.opts.Context.Value(hostKey{}).(host.Host)
	if !ok {
		err = fmt.Errorf("transport host shouldn't be nil")
		return
	}

	pid, ok := i.opts.Context.Value(protocolIDKey{}).(protocol.ID)
	if !ok {
		err = fmt.Errorf("transport protocolID shouldn't be nil")
		return
	}

	i.host = h
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

	var conn network.Conn

	defer func() {
		if err != nil {
			log.Errorf("transport %s dial peer %s error: %s", i.String(), addr, err)
		}
	}()

	if options.Timeout == 0 {
		options.Timeout = 5 * time.Second
	}

	// todo TLS

	conn, err = i.host.Network().DialPeer(options.Context, addr)
	if err != nil {
		err = fmt.Errorf("transport %s dial peer %s error: %s", i.String(), addr, err)
		return
	}

	c = &ipfsClient{
		conn: conn,
	}

	return
}

func (i *ipfsTransport) Listen(addr string, opts ...transport.ListenOption) (transport.Listener, error) {
	panic("implement me")
}

func (i *ipfsTransport) String() string {
	panic("implement me")
}
