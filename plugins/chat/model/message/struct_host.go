package message

// 	/ip4/${ip}/tcp/${port}/p2p/${id}
type host struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
	ID   string `json:"id"`
	// todo protocol
}
