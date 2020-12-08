package model

import (
	"net/http"
)

// todo Router 转移到合理位置
func Router() http.Handler {
	return router
}
