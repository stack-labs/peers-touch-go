package test

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"time"

	"github.com/joincloud/peers-touch-go/pubsub"
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	_ "github.com/joincloud/peers-touch-go/codec/json"
	_ "github.com/joincloud/peers-touch-go/logger/logrus"
	"github.com/joincloud/peers-touch-go/node"
)

func TestBrokerSub(t *testing.T) {
	/*
		mocknet := testingMockNet(ctx)
		node, clean := testingIPFSNode(ctx, t, mocknet)
		defer clean()

	ipfs := testingCoreAPI(t, node)*/
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Convey("test ", t, FailureHalts, func(c C) {
		h, _ := libp2p.New(ctx)
		n, err := node.NewNode(ctx, node.Host(h))
		_, _ = c.Printf("nodeId: %s\n", n.ID().String())
		_, _ = c.Printf("hid: %s\n", h.ID().String())

		c.So(err, ShouldBeNil)
		Convey("Create Broker", FailureHalts, func(c C) {
			ok := make(chan bool)
			Convey("Create Sub", FailureHalts, func(c C) {
				sub, err := n.Broker().Sub(ctx, "topic-test", func(event pubsub.Event) {
					_, _ = c.Printf("receive evt: %s-%s", event.Name(), event.Message())
					ok <- true
				})

				c.So(err, ShouldBeNil)
				c.So(sub, ShouldNotBeNil)
			})

			Convey("Create Pub", FailureHalts, func(c C) {
				err := n.Broker().Pub(ctx, pubsub.NewEvent("topic-test", pubsub.Message{Body: "Hello Everybody"}))
				c.So(err, ShouldBeNil)
			})
		})
	})

	select {
	case <-time.After(time.Second * 3):
		return
	}
}
