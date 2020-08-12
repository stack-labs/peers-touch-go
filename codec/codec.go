package codec

var (
	// todo, no Helper, more graceful
	Codecs = make(map[string]func(opts ...Option) Codec)
)

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	// name
	String() string
}

func NewCodec(name string, options ...Option) Codec {
	return Codecs[name](options...)
}
