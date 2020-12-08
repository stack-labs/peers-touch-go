package message

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/joincloud/home-platform/model/common"
	"github.com/joincloud/home-platform/util"
	log "github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/store"
)

func (m *Model) ConnectPeer(w http.ResponseWriter, r *http.Request) {
	// todo 一把梭，不用一个一个弄
	fromId := r.FormValue("fromId")
	toId := r.FormValue("toId")
	addr := r.FormValue("addr")
	name := r.FormValue("name")
	// time := r.FormValue("time")

	// todo 认证信息参数等
	p := Peer{
		Addr:   addr,
		ToID:   toId,
		FromID: fromId,
		Name:   name,
		Time:   time.Now().Unix(),
	}

	err := store.Write(
		p.ToRecord(),
		// todo 配置
		store.WriteTo(db, "peer_touched"),
	)
	rsp := common.Response{}
	if err != nil {
		rsp.Err = &common.Error{
			Error: err,
		}
	} else {
		rsp.Data = "ok"
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (m *Model) HandlerUpdatePeerAddr(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	toId := r.FormValue("toId")
	addr := r.FormValue("addr")
	peerAddrCache[bucketName(id, toId)] = addr

	log.Infof("register addr %s %s %s", id, toId, addr)

	util.RspJsonOK(w, "OK")
}
