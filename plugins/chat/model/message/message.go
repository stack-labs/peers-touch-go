package message

import (
	"bufio"
	"context"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/node"
	"github.com/joincloud/peers-touch-go/peer"
	"github.com/joincloud/peers-touch-go/pubsub"
	"github.com/joincloud/peers-touch-go/store"
)

var (
	peerAddrCache   = map[string]string{}
	msgCursorBuffer = newMessageCursorBuffer()
	db              = "join_cloud_node.db"
)

type Model struct {
	me          node.Node
	store       store.Store
	sendMsgLock sync.Mutex

	readerChan chan *Message
	writerChan chan *Message
	sendChan   chan *Message

	writerMap map[string]*writer
}

func (m *Model) Name() string {
	return "message"
}

func (m *Model) Hook(me node.Node) {
	m.me = me
}

func (m *Model) Handler(router *mux.Router) {
	// todo id 固定已经有用户
	router.HandleFunc("/pull-msg/{id:[0-9]+}/{from:[0-9]+}/{lastId:[0-9]+}", m.HandlerPullMsg).Methods("GET")
	router.HandleFunc("/push-msg/{id:[0-9]+}/{to:[0-9]+}", m.HandlerPushMsg).Methods("POST")
	router.HandleFunc("/connect-peer", m.ConnectPeer).Methods("POST")
	router.HandleFunc("/host-info", m.HostInfo).Methods("GET")
	router.HandleFunc("/update-peer-addr", m.HandlerUpdatePeerAddr).Methods("POST")
}

func (m *Model) Store(store store.Store) {
	m.store = store
}

func (m *Model) Start() (err error) {
	m.readerChan = make(chan *Message)
	m.writerChan = make(chan *Message)
	m.sendChan = make(chan *Message)
	m.writerMap = map[string]*writer{}

	go m.waitChanStream()
	go m.listenReaderChan()
	go m.listenSendChan()
	// todo 加载最新的消息id
	// todo 初始化Buffer
	m.loadMsgCursorBuffer()
	// todo 初始化Peer推送管道，见../peer模块
	// m.connectToPeersForChat()

	return
}

func (m *Model) listenReaderChan() {
	for {
		select {
		case msg := <-m.readerChan:
			_, err := m.saveMsg(msg)
			if err != nil {
				log.Errorf("listenReaderChan save msg err: %s", err)
			}
		}
	}
}

func (m *Model) listenSendChan() {
	for {
		select {
		case msg := <-m.sendChan:
			m.sendMsg(msg)
		}
	}
}

func (m *Model) waitChanStream() {
	// todo error
	_ = m.me.Broker().Touch(context.TODO(),
		pubsub.TouchStreamHandler(func(s peer.Stream) {
			rmtAddr := s.Conn().RemoteMultiaddr().String()
			rmtID := s.Conn().RemotePeer().Pretty()

			log.Infof("got new stream connection from %s and its id is %s, metadata: %v", rmtAddr, rmtID, s.Stat().Extra)
			rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

			// 创建桶
			w := newWriter("msg_1_2", m.writerChan)
			r := newReader(m.readerChan)

			go w.Writer(rw)
			go r.Reader(rw)

			peerAddrCache[w.bucketName] = rmtAddr
			m.writerMap[w.bucketName] = w
		}),
	)
}

func (m *Model) loadMsgCursorBuffer() {
	// todo 查出有多少条，可以建个key维护，先查这个key，再看是不是大，不大就一次性load出来
	records, err := m.store.List(store.ListFrom(db, "msg_cursor"),
		// 假设只有20条
		store.ListLimit(20),
		store.ListOffset(0),
	)
	if err != nil {
		log.Errorf("load msg cursor buffer err: %s", err)
		return
	}

	for _, record := range records {
		msg := recordToMsg(record)
		log.Infof("load msg cursor for %v", msg)
		msgCursorBuffer.push(msg)
	}
}

func (m *Model) connectToPeersForChat() {
	// todo 查出peers-chat桶中所有字段
	// err 处理
	records, err := m.store.List(store.ListFrom("", "peers_info"),
		// todo 合理处理
		store.ListLimit(10),
		store.ListOffset(0),
	)
	if err != nil {
		log.Errorf("", err)
		return
	}

	// todo 字段解析映射
	for _, record := range records {
		peer := &Peer{}
		_ = peer.FromRecord(record)
		_, err := m.connect(peer.key(), peer.FromID)
		if err != nil {
			log.Errorf("connect to peer error: %s", err)
		}
	}
}
