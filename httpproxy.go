package httpproxy

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/499689317/go-log"
	"github.com/gin-gonic/gin"
)

type Configurable interface {
	HTTPListenAddr() string
	HTTPTimeout() time.Duration
}

type Server struct {
	config  Configurable
	server  *http.Server
	handler *gin.Engine
}

func NewServer(c Configurable) *Server {

	gin.SetMode(gin.ReleaseMode)

	// g := gin.Default()
	g := gin.New()

	g.Use(gin.ErrorLogger())
	g.Use(gin.Recovery())

	s := &Server{
		server: &http.Server{
			Addr:         c.HTTPListenAddr(),
			Handler:      g,
			ReadTimeout:  c.HTTPTimeout() * time.Second,
			WriteTimeout: c.HTTPTimeout() * time.Second,
		},
	}

	// TODO
	s.config = c
	s.handler = g

	Init(g)

	log.Info().Str("Listen At", c.HTTPListenAddr()).Msg("New HTTP Server ok")
	return s
}

func (s *Server) Start() {

	if s.server == nil {
		log.Error().Msg("New HTTP Server failed")
		return
	}

	s.server.ListenAndServe()
}

func (s *Server) Run() {

	go s.Start()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if e := s.server.Shutdown(ctx); e != nil {
		log.Fatal().Err(e).Msg("Server Shutdown:")
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info().Msg("timeout of 5 seconds.")
	}
	log.Info().Msg("Server exiting")
}
