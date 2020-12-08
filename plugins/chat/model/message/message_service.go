package message

import (
	"context"
	"fmt"
	"time"

	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/pubsub"
	"github.com/joincloud/peers-touch-go/store"
)

func (m *Model) sendMsg(msg *Message) {
	m.sendMsgLock.Lock()
	defer m.sendMsgLock.Unlock()

	if w, ok := m.writerMap[msg.bucketName]; ok {
		w.sendChan <- msg
	} else { // 没有则创建连接
		w, err := m.connect(msg.ID, msg.TargetID)
		if err != nil {
			log.Errorf("sendMsg after connecting peer error: %s", err)
			return
		}

		w.sendChan <- msg
	}
}

func (m *Model) connect(id, targetID string) (w *writer, err error) {
	bucketName := bucketName(id, targetID)
	// 找到to的地址
	addr, ok := peerAddrCache[bucketName]
	if !ok {
		err = fmt.Errorf("cant find peer addr for %s", bucketName)
		return
	}

	if w, ok := m.writerMap[bucketName]; ok {
		return w, nil
	}

	w = newWriter(bucketName, m.writerChan)
	r := newReader(m.readerChan)

	// todo error
	_ = m.me.Broker().Touch(context.TODO(),
		pubsub.TouchReader(r.Reader),
		pubsub.TouchWriter(w.Writer),
		pubsub.TouchAddr(addr),
	)
	peer := &Peer{
		Name:   bucketName,
		ToID:   targetID,
		FromID: id,
		Addr:   addr,
		Time:   time.Now().Unix(),
	}

	_ = m.store.Write(peer.ToRecord(), store.WriteTo("", "peers_info"))

	m.writerMap[bucketName] = w

	return
}

func (m *Model) saveMsg(msg *Message) (ret *Message, err error) {
	defer func() {
		if err != nil {
			log.Errorf("send msg error: %s", err)
		}
	}()

	if msg.bucketName == "" {
		msg.bucketName = bucketName(msg.ID, msg.TargetID)
	}

	// 找到上一条消息
	msgTmp := msgCursorBuffer.getLastId(msg.ID, msg.TargetID)
	if msgTmp == nil {
		// todo 不锁方法，要锁id,to资源
		m.sendMsgLock.Lock()
		defer m.sendMsgLock.Unlock()

		// 创世消息
		genesisMsg := newGenesisBlock(msg.ID, msg.TargetID)
		// todo handler error
		_ = m.store.Write(
			msgToRecord(genesisMsg),
			// todo 配置
			store.WriteTo(db, msg.bucketName))

		msg = newMessage(plus1(genesisMsg.Idx), msg.ID, msg.TargetID, msg.Data, genesisMsg.Hash)
	} else {
		log.Infof("lastId is %s", msgTmp.Idx)
		msg = newMessage(plus1(msgTmp.Idx), msg.ID, msg.TargetID, msg.Data, msg.Hash)
	}

	// 更新缓存
	// todo 优化发送逻辑
	// cache the last Id
	// 先入库
	rc := msgToRecord(msg)
	err = m.store.Write(
		rc,
		// todo 配置
		store.WriteTo(db, msg.bucketName))
	if err != nil {
		err = fmt.Errorf("store write msg err: %s", err)
		return
	}

	err = m.store.Write(
		&store.Record{
			Key:   fmt.Sprintf("%s_%s", msg.ID, msg.TargetID),
			Value: msg.Serialize(),
		},
		// todo 配置
		store.WriteTo(db, "msg_cursor"))
	if err != nil {
		err = fmt.Errorf("store write msg err: %s", err)
		return
	}

	msgCursorBuffer.push(*msg)

	return msg, nil
}

func bucketName(id, targetID string) string {
	return fmt.Sprintf("msg_%s_%s", id, targetID)
}
