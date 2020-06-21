package registry

import (
	"context"
	"time"
)

type Options struct {
}

type NodeOptions struct {
	Name string `json:"name"`
}

type RegisterOptions struct {
	TTL     time.Duration
	Context context.Context
}
