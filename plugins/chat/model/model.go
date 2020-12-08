package model

import (
	"context"
	"crypto/rand"
	"fmt"
	"sync"

	"github.com/gorilla/mux"
	"github.com/joincloud/home-platform/model/message"
	"github.com/joincloud/home-platform/model/ping"
	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/node"
	"github.com/joincloud/peers-touch-go/store"
	"github.com/joincloud/peers-touch-go/store/bolt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
)

var (
	db     = "join_cloud_node.db"
	me     node.Node
	st     store.Store
	router = &mux.Router{}
	mut    sync.Mutex
)

type Model interface {
	Hook(n node.Node)
	Handler(router *mux.Router)
	Store(store store.Store)
	Start() (err error)
	Name() string
}

func Hook(m Model) {
	mut.Lock()
	defer mut.Unlock()

	m.Hook(me)
	m.Store(st)
	m.Handler(router)
	if err := m.Start(); err != nil {
		log.Errorf("start model %s err: %s", m.Name(), err)
	}
}

func Init(httpPort, ipfsPort int, dbDir string) {
	log.Infof("init store")
	var err error

	if dbDir != "" {
		db = dbDir + db
	}

	st, err = initStore()
	if err != nil {
		panic(err)
	}

	// todo prepare host
	r := rand.Reader
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		panic(err)
	}

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", ipfsPort))
	host, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)

	if err != nil {
		// 严重灾难，直接宕机
		// todo 更好的处理方式
		panic(err)
	}

	// 打印调试信息
	log.Infof("my ipfs addr: '/ip4/127.0.0.1/tcp/%v/p2p/%s'", ipfsPort, host.ID().Pretty())

	me, _ = node.NewNode(
		// todo context 管理
		context.TODO(),
		node.Host(host),
	)

	log.Infof("init modes")

	pi := &ping.Model{}
	Hook(pi)
	msg := &message.Model{}
	Hook(msg)
}

func initStore() (st store.Store, err error) {
	// 一个数据库
	// 一个库N个桶
	// 一个聊天对象一个桶，相当于用户互相之间有一个文件分片
	// eg. A与B聊天，信息会放在 MSG_A_B 桶下面
	// 每条信息以K-V信息存储
	// K为序号：$idx
	// V为聊天内容+元数据（Hash，时间，IP等等附属信息）
	st, err = bolt.NewStore(
		// 先硬编码
		store.Database(db),
		// todo 统一全局context
		store.WithContext(context.TODO()),
		bolt.CreateBuckets([]string{
			"msg_1_2",    // 消息存储桶
			"msg_2_1",    // 消息存储桶
			"peers_info", // 端-端互联聊天配置桶
			"msg_cursor", // 消息游标桶
		}), // todo 加上动态创建桶
	)

	store.Register(st, true)

	if err != nil {
		log.Errorf("node inits store error: %s", err)
		return nil, err
	}

	log.Infof("home-platform store inited")
	return
}
