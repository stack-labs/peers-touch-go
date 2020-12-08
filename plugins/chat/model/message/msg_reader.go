package message

import (
	"bufio"
	"sync"

	log "github.com/joincloud/peers-touch-go/logger"
)

var (
	rwLock sync.Mutex
)

type reader struct {
	readerChan chan *Message
}

func (r *reader) Reader(rw *bufio.ReadWriter) {
	for {
		bs, err := rw.ReadBytes(byte('\n'))
		if bs == nil {
			return
		}

		if err != nil {
			log.Errorf("model reader read msg error: %s", err)
			return
		}

		if len(bs) > 0 {
			log.Debugf("model reader received msg: %v", string(bs))

			msg := &Message{}
			err = msg.Deserialize(bs)
			if err != nil {
				log.Errorf("model reader deserialize msg error: %s", err)
				return
			}
			// set bucketKey
			msg.bucketName = bucketName(msg.TargetID, msg.ID)
			r.readerChan <- msg
		}
	}
}

func newReader(readerChan chan *Message) *reader {
	return &reader{
		readerChan: readerChan,
	}
}
