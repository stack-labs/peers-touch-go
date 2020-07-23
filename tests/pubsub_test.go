package test

import (
	"context"
	"testing"

	"github.com/joincloud/peers-touch/pubsub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBrokerSub(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mocknet := testingMockNet(ctx)
	node, clean := testingIPFSNode(ctx, t, mocknet)
	defer clean()

	ipfs := testingCoreAPI(t, node)

	Convey("test ", func(c C) {
		Convey("Create Broker", func(c C) {
			b := pubsub.NewBroker(pubsub.BrokerCoreAPI(ipfs))
			c.So(b, ShouldNotBeNil)

			Convey("Create Sub", func(c C) {
				sub, err := b.Sub(ctx, "topic-test", func(event pubsub.Event) {
					c.Printf("receive evt: %s-%s", event.Topic(), event.Message())
				})

				c.So(err, ShouldBeNil)
				c.So(sub, ShouldNotBeNil)
			})
		})
	})

}
