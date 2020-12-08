package message

import (
	"fmt"
	"sync"
)

type messageCursorBuffer struct {
	cache map[string]Message
	sync.Mutex
}

func (m *messageCursorBuffer) push(msg Message) {
	m.Lock()
	defer m.Unlock()

	m.cache[m.key(msg.ID, msg.TargetID)] = msg
}

func (m *messageCursorBuffer) getLastId(id, from string) *Message {
	m.Lock()
	defer m.Unlock()

	if msg, ok := m.cache[m.key(id, from)]; ok {
		return &msg
	}

	return nil
}

func (m *messageCursorBuffer) key(id, from string) string {
	return fmt.Sprintf("%s_%s", id, from)
}

func newMessageCursorBuffer() *messageCursorBuffer {
	return &messageCursorBuffer{
		cache: map[string]Message{},
	}
}
