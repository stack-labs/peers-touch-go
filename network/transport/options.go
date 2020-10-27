package transport

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/joincloud/peers-touch-go/codec"
)

type Options struct {
	Addrs     []string
	Codec     codec.Codec
	Secure    bool
	TLSConfig *tls.Config
	Timeout   time.Duration
	Context   context.Context
}

type DialOptions struct {
	Stream  bool
	Timeout time.Duration
	Context context.Context
}

type ListenOptions struct {
	Context context.Context
}

func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Codec(c codec.Codec) Option {
	return func(o *Options) {
		o.Codec = c
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

func Secure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

func Stream() DialOption {
	return func(o *DialOptions) {
		o.Stream = true
	}
}

func DialTimeout(d time.Duration) DialOption {
	return func(o *DialOptions) {
		o.Timeout = d
	}
}
