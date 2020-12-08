package ping

import (
	"github.com/gorilla/mux"
	"github.com/joincloud/peers-touch-go/node"
	"github.com/joincloud/peers-touch-go/store"
)

type Model struct {
}

func (m *Model) Hook(me node.Node) {
	// do nothing
}

func (m *Model) Handler(router *mux.Router) {
	router.HandleFunc("/ping", Ping).Methods("GET")
}

func (m *Model) Start() error {
	// do nothing

	return nil
}

func (m *Model) Store(store.Store) {
	return
}

func (m *Model) Name() string {
	return "ping"
}
