package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/stack-labs/peers-touch-go/platform"
)

var (
	mux sync.Mutex
)

type Platform struct {
	client      *http.Client
	initOpts    *platform.InitOptions
	connectOpts *platform.ConnectOptions
	URL         string
}

func (p *Platform) Init(opts ...platform.InitOption) error {
	mux.Lock()
	defer mux.Unlock()

	p.client = &http.Client{}

	return nil
}

func (p *Platform) Connect(opts ...platform.ConnectOption) error {
	nodeJSON, _ := json.Marshal(p.initOpts.Node)
	req, err := http.NewRequest("POST", p.URL, bytes.NewBuffer(nodeJSON))
	if err != nil {
		return fmt.Errorf("")
	}

	resp, err := p.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return nil
}

func (p *Platform) Close() error {
	panic("implement me")
}
