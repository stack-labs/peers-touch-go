package message

import (
	"encoding/json"
	"fmt"

	"github.com/joincloud/peers-touch-go/store"
)

// Peer is the node we need to connect
// todo block
type Peer struct {
	// the user's id in connection source node
	FromID string `json:"fromId,omitempty"`
	// the user's id in connection dest node
	ToID string `json:"toId,omitempty"`
	// dest node addr
	Addr string `json:"addr"`
	// Peer's nikename
	Name string `json:"name"`
	// peer first touched time
	Time int64 `json:"time"`
	// todo RSA
}

func (p *Peer) ToRecord() *store.Record {
	sr := &store.Record{}
	sr.Key = p.key()
	sr.Value = p.Bytes()

	return sr
}

func (p *Peer) FromRecord(sr *store.Record) (err error) {
	// todo check sr
	// todo use codec
	temp := &Peer{}
	err = json.Unmarshal(sr.Value, temp)
	if err != nil {
		return
	}

	p.Name = temp.Name
	p.ToID = temp.ToID
	p.FromID = temp.FromID
	p.Addr = temp.Addr
	p.Time = temp.Time

	return
}

func (p *Peer) Bytes() []byte {
	// todo use codec
	bytes, _ := json.Marshal(p)
	return bytes
}

func (p *Peer) key() string {
	return fmt.Sprintf("%s_%s_%s", p.FromID, p.ToID, p.Name)
}
