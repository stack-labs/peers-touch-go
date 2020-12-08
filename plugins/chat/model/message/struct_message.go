package message

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/store"
)

type Message struct {
	Idx           string `json:"idx"`
	ID            string `json:"id"`
	TargetID      string `json:"targetId"`
	Timestamp     int64  `json:"timestamp"`
	Data          string `json:"data"`
	PrevBlockHash []byte `json:"-"`
	Hash          []byte `json:"-"` // todo 优化返回
	Nonce         int    `json:"-"` // todo 种子记下

	bucketName string // 不作为Hash计算，只是用来找表
}

func (b *Message) Serialize() []byte {
	// todo use codec
	bs, _ := json.Marshal(b)
	return bs
}

func (b *Message) Deserialize(bs []byte) error {
	tmp := &Message{}
	_ = json.Unmarshal(bs, tmp)

	b.Idx = tmp.Idx
	b.ID = tmp.ID
	b.TargetID = tmp.TargetID
	b.Timestamp = tmp.Timestamp
	b.Data = tmp.Data
	b.PrevBlockHash = tmp.PrevBlockHash
	b.Hash = tmp.Hash
	b.Nonce = tmp.Nonce

	return nil
}

func msgToRecord(msg *Message) *store.Record {
	sr := &store.Record{
		Key: msg.Idx,
	}

	log.Debugf("key in msg to record is: %s", sr.Key)

	sr.Value, _ = json.Marshal(msg)
	return sr
}

func recordToMsg(sr *store.Record) Message {
	msg := Message{}
	_ = json.Unmarshal(sr.Value, &msg)

	return msg
}

func newMessage(idx, id, targetID, data string, prevBlockHash []byte) *Message {
	// todo idx
	block := &Message{
		Idx:           idx,
		ID:            id,
		TargetID:      targetID,
		Data:          data,
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Nonce:         0,

		bucketName: bucketName(id, targetID),
	}
	pow := newProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func newGenesisBlock(id, targetID string) *Message {
	return newMessage(fmtIdx(0), id, targetID, "Genesis Message", []byte{})
}

func fmtIdx(idx int) string {
	return fmt.Sprintf("%08d", idx)
}

func plus1(idx string) string {
	i, _ := strconv.Atoi(idx)
	i++
	return fmtIdx(i)
}
