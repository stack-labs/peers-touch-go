package peers

func NewPeer(opts ...Option) Peer {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	p := newPeer(o)

	return p
}
