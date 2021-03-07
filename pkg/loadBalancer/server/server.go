package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Config struct {
	Address string
	Router  *gin.Engine
}

type Server struct {
	server *http.Server
	Router *gin.Engine
}

func New(config *Config) *Server {
	return &Server{
		server: &http.Server{
			Addr:         config.Address,
			Handler:      config.Router,
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 60 * time.Second,
		},
		Router: config.Router,
	}
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}