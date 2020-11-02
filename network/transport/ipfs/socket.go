package ipfs

import (
	"bufio"
	"fmt"
	"github.com/joincloud/peers-touch-go/codec"
	log "github.com/joincloud/peers-touch-go/logger"

	"github.com/joincloud/peers-touch-go/network/transport"
	"github.com/libp2p/go-libp2p-core/network"
)

type ipfsTransportSocket struct {
	stream network.Stream
	w      *bufio.Writer
	r      *bufio.Reader
	codec  codec.Codec
	local  string
	remote string
}

func (i *ipfsTransportSocket) Recv(msg *transport.Message) error {
	return nil
}

func (i *ipfsTransportSocket) Send(msg *transport.Message) (err error) {
	if msg == nil {
		return nil
	}

	defer func() {
		if err != nil {
			log.Infof("ipfs socket send bytes error: %s", err)
		}
	}()

	bytes, err := i.codec.Marshal(msg)
	if err != nil {
		err = fmt.Errorf("ipfs socket encode msg error: %s", err)
		return
	}

	_, err = i.w.Write(bytes)
	if err != nil {
		err = fmt.Errorf("ipfs socket encode msg error: %s", err)
		return
	}

	err = i.w.WriteByte(byte('\n'))
	if err != nil {
		err = fmt.Errorf("ipfs socket write bytes error: %s", err)
		return
	}

	err = i.w.Flush()
	if err != nil {
		err = fmt.Errorf("ipfs socket flush writer error: %s", err)
		return
	}

	return
}

func (i *ipfsTransportSocket) Close() error {
	return i.stream.Close()
}

func (i *ipfsTransportSocket) Local() string {
	return i.local
}

func (i *ipfsTransportSocket) Remote() string {
	return i.remote
}
