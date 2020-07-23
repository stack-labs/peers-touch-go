package test

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBrokerSub(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mocknet := testingMockNet(ctx)
	node, clean := testingIPFSNode(ctx, t, mocknet)
	defer clean()

	ipfs := testingCoreAPI(t, node)

}
