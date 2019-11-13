# go-http

+ install go-http

`go get -u github.com/499689317/go-http`


+ how to use it

````

import (
	"time"

	"github.com/499689317/go-http"
	"github.com/499689317/go-log"
	zlog "github.com/rs/zerolog/log"
)

type conf struct {
	HttpListenAddr string
	HttpTimeout    time.Duration
}
func (c *conf) HTTPListenAddr() string {
	return c.HttpListenAddr
}
func (c *conf) HTTPTimeout() time.Duration {
	return c.HttpTimeout
}

func main() {

	zlog.Logger = zlog.With().Caller().Logger()
	log.Init()
	log.SetLogLevel(1)

	// conf
	c := &conf{
		HttpListenAddr: ":8060",
		HttpTimeout:    30
	}

	// http
	h := httpproxy.NewServer(c)
	go h.Start()
	select{}

	// or
	// h.Run()
}

````