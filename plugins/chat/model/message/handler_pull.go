package message

import (
	"github.com/joincloud/home-platform/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joincloud/peers-touch-go/store"
)

func (m *Model) HandlerPullMsg(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fromId := mux.Vars(r)["from"]
	lastMsgId, _ := strconv.Atoi(mux.Vars(r)["lastId"])

	records, err := store.List(
		store.ListFrom("", bucketName(id, fromId)),
		store.ListLimit(10),
		store.ListOffset(uint(lastMsgId)),
	)

	if err != nil {
		util.RspJsonErr(w, util.Error{Msg: err})
		return
	} else {
		var ret []interface{}
		for _, rc := range records {
			ret = append(ret, recordToMsg(rc))
		}
		util.RspJsonOK(w, ret)
		return
	}
}
