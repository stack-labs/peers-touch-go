package test

import (
	"context"
	"testing"
	"time"

	"github.com/joincloud/peers-touch-go/pubsub"
	. "github.com/smartystreets/goconvey/convey"

	_ "github.com/joincloud/peers-touch-go/codec/json"
	_ "github.com/joincloud/peers-touch-go/logger/logrus"
)

func TestBrokerSub(t *testing.T) {
	ctx := context.Background()

	mocknet := testingMockNet(ctx)
	node, clean := testingIPFSNode(ctx, t, mocknet)
	defer clean()

	ipfs := testingCoreAPI(t, node)

	Convey("test ", t, FailureHalts, func(c C) {
		Convey("Create Broker", FailureHalts, func(c C) {
			b := pubsub.NewBroker(pubsub.BrokerCoreAPI(ipfs))
			c.So(b, ShouldNotBeNil)

			ok := make(chan bool)
			Convey("Create Sub", FailureHalts, func(c C) {
				sub, err := b.Sub(ctx, "topic-test", func(event pubsub.Event) {
					_, _ = c.Printf("receive evt: %s-%s", event.Name(), event.Message())
					ok <- true
				})

				c.So(err, ShouldBeNil)
				c.So(sub, ShouldNotBeNil)
			})

			Convey("Create Pub", FailureHalts, func(c C) {
				err := b.Pub(ctx, pubsub.NewEvent("topic-test", pubsub.Message{Body: "Hello Everybody"}))
				c.So(err, ShouldBeNil)
			})

			select {
			case <-ok:
				return
			case <-time.After(time.Second * 4):
				return
			}
		})
	})
}
