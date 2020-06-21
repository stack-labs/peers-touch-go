package registry

type Node struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	NameSpace string            `json:"namespace"`
	Address   Address           `json:"address"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []*Endpoint       `json:"endpoints"`
}

type Address struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}
