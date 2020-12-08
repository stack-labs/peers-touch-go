package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/joincloud/home-platform/model"
	log "github.com/joincloud/peers-touch-go/logger"

	_ "github.com/joincloud/peers-touch-go/codec/json"
	_ "github.com/joincloud/peers-touch-go/logger/logrus"
)

func main() {
	// 初始化组件
	httpPort := flag.Int("http-port", 8090, "Source port number")
	ipfsPort := flag.Int("ipfs-port", 8091, "Source port number")
	dbDir := flag.String("db-dir", "", "db dir")
	flag.Parse()

	model.Init(*httpPort, *ipfsPort, *dbDir)

	srv := &http.Server{
		Handler:      model.Router(),
		Addr:         fmt.Sprintf(":%d", *httpPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("home-platform starts OK!")
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
