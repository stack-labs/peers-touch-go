package registry

import (
	"context"
	"fmt"

	"github.com/joincloud/peers-touch－go/peer"
	peer2 "github.com/libp2p/go-libp2p-core/peer"
)

type Node struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	NameSpace string            `json:"namespace"`
	Address   Address           `json:"address"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []*Endpoint       `json:"endpoints"`

	// todo
	Host peer.Host
}

func (n *Node) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.Host.Connect(context.Background(), peer2.AddrInfo(pi))
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

type Address struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}
