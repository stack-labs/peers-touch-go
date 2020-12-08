package ping

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}
