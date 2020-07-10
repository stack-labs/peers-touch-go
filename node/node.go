package node

import (
	"sync"

	"github.com/ipfs/go-ipfs/core/coreapi"
)

type node struct {
	ipfs coreapi.CoreAPI

	muPubSub sync.RWMutex
	muIPFS   sync.RWMutex
}

func (n *node) IPFS() coreapi.CoreAPI {
	n.muIPFS.RLock()
	defer n.muIPFS.RUnlock()

	return n.ipfs
}
