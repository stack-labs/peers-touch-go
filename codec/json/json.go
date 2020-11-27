package json

import (
	"encoding/json"

	"github.com/stack-labs/peers-touch-go/codec"
)

type Options struct {
	Name string
}

type Codec struct {
	name string
}

func init() {
	codec.Codecs["json"] = func(options ...codec.Option) codec.Codec {
		return NewCodec(options...)
	}
}

func (c *Codec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *Codec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (c *Codec) String() string {
	return c.name
}

func (c *Codec) New(options ...codec.Option) codec.Codec {
	if options == nil {
		return c
	}

	return NewCodec(options...)
}

func NewCodec(options ...codec.Option) codec.Codec {
	opts := &Options{}
	for _, option := range options {
		option(opts)
	}

	if opts.Name == "" {
		opts.Name = "json"
	}

	return &Codec{
		name: opts.Name,
	}
}
