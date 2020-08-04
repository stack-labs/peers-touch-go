package json

import (
	"encoding/json"

	"github.com/joincloud/peers-touch-go/codec"
)

type Options struct {
}

type Codec struct {
}

func (c *Codec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *Codec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (c *Codec) String() string {
	return "json"
}

func NewCodec(options ...codec.Option) codec.Codec {
	opts := &Options{}
	for _, option := range options {
		option(opts)
	}
	return &Codec{}
}
