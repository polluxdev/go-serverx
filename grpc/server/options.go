package server

import (
	"fmt"
	"time"
)

// Option -.
type Option func(*Server)

// Port -.
func Port(addr string) Option {
	return func(s *Server) {
		s.addr = fmt.Sprintf(":%s", addr)
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
