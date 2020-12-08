package message

import (
	"bufio"

	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/store"
)

type writer struct {
	bucketName string
	sendChan   chan *Message
	store      store.Store
}

func (w *writer) Writer(rw *bufio.ReadWriter) {
	for {
		select {
		case msg := <-w.sendChan:
			{
				if msg.bucketName == w.bucketName {
					_, _ = rw.Write(msg.Serialize())
					_ = rw.WriteByte(byte('\n'))
					_ = rw.Flush()
					continue
				}

				log.Debugf("%s ignore the msg of %s", w.bucketName, msg.bucketName)
			}
		}
	}
}

func newWriter(key string, sendChan chan *Message) (rwRet *writer) {
	rwLock.Lock()
	defer rwLock.Unlock()

	rw := &writer{
		bucketName: key,
		sendChan:   sendChan,
	}

	return rw
}
