package message

import (
	"encoding/json"
	"net/http"

	"github.com/joincloud/home-platform/model/common"
)

func (m *Model) HostInfo(w http.ResponseWriter, r *http.Request) {
	rsp := common.Response{
		Data: &host{
			IP:   "127.0.0.1",
			Port: "8090",
			ID:   m.me.Options().Host.ID().Pretty(),
		},
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
