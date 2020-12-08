package common

type Error struct {
	Error error  `json:"error,omitempty"`
	Code  string `json:"code,omitempty"`
}

type Response struct {
	Err  *Error      `json:"error,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
