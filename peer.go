package peers

type Peer interface {
	Init(...Options) error
	Run() error
}

type peer struct {
	options *Options
}

func (p *peer) Init(options ...Options) error {
	return nil
}

func (p *peer) Run() error {
	return nil
}

func newPeer(options *Options) Peer {
	return &peer{
		options: options,
	}
}
