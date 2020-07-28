package registry

import (
	"context"
	"github.com/libp2p/go-libp2p-core/host"
	"time"
)

type Options struct {
}

type NodeOptions struct {
	Name string `json:"name"`
}

type RegisterOptions struct {
	TTL      time.Duration
	Context  context.Context
	Host     host.Host
	Interval time.Duration
}
