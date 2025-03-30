package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	_defaultAddr            = ":8082"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	addr            string
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler: handler,
		Addr:    _defaultAddr,
	}

	s := &Server{
		addr:            _defaultAddr,
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Start -.
func (s *Server) Start() {
	log.Printf("[websocket] server is starting on %s...", s.addr)

	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
