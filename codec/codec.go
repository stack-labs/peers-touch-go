package codec

var (
	Codecs = make(map[string]Codec)
)

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	// name
	String() string
}

func NewCodec(name string, opts ...Options) Codec {
	return
}
