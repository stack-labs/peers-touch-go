package test

import (
	"context"
	"testing"

	"github.com/joincloud/peers-touch－go/pubsub"
	. "github.com/smartystreets/goconvey/convey"
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

			Convey("Create Sub", FailureHalts, func(c C) {
				sub, err := b.Sub(ctx, "topic-test", func(event pubsub.Event) {
					c.Printf("receive evt: %s-%s", event.Topic(), event.Message())
				})

				c.So(err, ShouldBeNil)
				c.So(sub, ShouldNotBeNil)
			})

			Convey("Create Pub", FailureHalts, func(c C) {
				err := b.Pub(ctx, pubsub.NewEvent("topic-test", []byte("Hello")))
				c.So(err, ShouldBeNil)
			})
		})
	})
}
