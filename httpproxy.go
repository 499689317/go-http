package httpproxy

import (
	"net/http"
	"time"

	"github.com/499689317/go-log"
	"github.com/gin-gonic/gin"
)

type Configurable interface {
	HTTPListenAddr() string
	HTTPTimeout() time.Duration
}

type Server struct {
	config Configurable
	server *http.Server
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

	s.Start()

	select {}
}
