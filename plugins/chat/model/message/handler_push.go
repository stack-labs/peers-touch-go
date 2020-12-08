package message

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joincloud/home-platform/util"
)

/*
func PushMsg(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	to := mux.Vars(r)["to"]
	str := r.FormValue("msg")
	msg := str2msg(str)

	// todo check params

	// cache the last Id
	err := store.Write(
		msgToRecord(msg),
		// todo 配置
		store.WriteTo("join_cloud_node.db", fmt.Sprintf("msg_%s_%s", id, to)))
	// buffer refresh
	key := "msg_" + strconv.Itoa(id) + "_" + strconv.Itoa(to)
	//如果字典中没有对应的key, 那么就是增加
	//如果字典中已经有对应的key, 那么就是修改
	buf[key] = msg
	//下面是更新表msg_cursor
	db, err := bolt.Open("join_cloud_node.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	b := tx.Bucket([]byte("msg_cursor"))
	//每次有新消息到来时，更新表
	err := b.Put(msg.Idx, key)
	if err != nil {
		log.Fatal(err)
	}
	//写完记得关闭表和库
	defer tx.Rollback()
	defer db.Close()

	rsp := common.Response{}
	if err != nil {
		rsp.Err = &common.Error{
			Error: err,
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}*/

//func msgToRecord(msg string) *store.Record {
//	sr := &store.Record{
//		Key: "1",
//	}
//	data := newMessage(msg, []byte("1")) // todo 上一消息的Hash
//	sr.Value, _ = json.Marshal(data)
//
//	return sr
//}
/*
func str2msg(str string) Message {
	db, err := bolt.Open("join_cloud_node.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	b3 := tx.Bucket([]byte("msg_cursor"))
	//获取桶标签
	key := "msg_" + strconv.Itoa(id) + "_" + strconv.Itoa(to)
	//获取桶中最大id
	value := b3.Get(key)
	//根据id从db中获取record
	b1 := tx.Bucket([]byte("msg_1_2"))
	//根据id从表中获取record
	prevRecord := b1.Get(value)
	//将prevRecord反序列化转成Message，并获取hash值
	prevMessage := recordToMsg(prevRecord)
	_Hash := prevMessage.Hash
	// not change str
	msg := newMessage(str, _Hash)
	//msg := newMessage(str, []byte("1"))// todo 上一消息的Hash
	return msg
}*/

func (m *Model) HandlerPushMsg(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	to := mux.Vars(r)["to"]
	data := r.FormValue("msg")

	// todo check params
	msg := &Message{
		ID:         id,
		TargetID:   to,
		Data:       data,
		bucketName: bucketName(id, to),
	}
	msg, err := m.saveMsg(msg)
	if err != nil {
		util.RspJsonErr(w, util.Error{Msg: err})
		return
	}

	m.sendChan <- msg

	util.RspJsonOK(w, "OK")
	return
}
