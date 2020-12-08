package broker

import "github.com/joincloud/peers-touch-go/pubsub"

type User struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	RemarkName string   `json:"remarkName,omitempty"`
	Tags       []string `json:"tags,omitempty"` // 用户标签
}

type Message struct {
	m    pubsub.Message
	ID   int    `json:"id"`
	Typ  int    `json:"typ"`  // 消息类型，1 普通消息，2 群消息
	User User   `json:"user"` // 来源用户
	Body string `json:"body"` // 消息内容，todo 支持更多结构内容
}
