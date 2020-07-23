package test

import (
	"context"
	"testing"

	ipfsCore "github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	mock "github.com/ipfs/go-ipfs/core/mock"
	iface "github.com/ipfs/interface-go-ipfs-core"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	. "github.com/smartystreets/goconvey/convey"
)

type cleanFunc func()

func newIPFSAPI(t *testing.T, core *ipfsCore.IpfsNode) (api iface.CoreAPI) {
	t.Helper()

	var err error
	if api, err = coreapi.NewCoreAPI(core); err != nil {
		t.Fatal(err)
	}

	Convey("")

	return
}

func testingMockNet(ctx context.Context) mocknet.Mocknet {
	return mocknet.New(ctx)
}

// MakeIPFS Creates a new IPFS node for testing purposes
func testingIPFSNode(ctx context.Context, t *testing.T, m mocknet.Mocknet) (*ipfsCore.IpfsNode, cleanFunc) {
	t.Helper()

	core, err := ipfsCore.NewNode(ctx, &ipfsCore.BuildCfg{
		Online: true,
		Host:   mock.MockHostOption(m),
		ExtraOpts: map[string]bool{
			"pubsub": true,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	return core, func() {
		core.Close()
	}
}

func testingCoreAPI(t *testing.T, core *ipfsCore.IpfsNode) (api iface.CoreAPI) {
	t.Helper()

	var err error
	if api, err = coreapi.NewCoreAPI(core); err != nil {
		t.Fatal(err)
	}

	return
}
