package httpproxy

import (
	"testing"
	"time"
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

/*
func TestHttpServerStart(t *testing.T) {

	c := &conf{
		HttpListenAddr: ":8060",
		HttpTimeout: 30,
	}

	s := NewServer(c)

	go s.Start()

	// todo some thing.....


	select{}
}
*/

func TestHttpServerRun(t *testing.T) {

	c := &conf{
		HttpListenAddr: ":8070",
		HttpTimeout:    30,
	}

	s := NewServer(c)

	s.Run()
}
