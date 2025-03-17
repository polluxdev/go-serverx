package fasthttp

import (
	"time"

	"github.com/valyala/fasthttp"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *fasthttp.Server
	notify          chan error
	shutdownTimeout time.Duration
	addr            string
}

// New -.
func New(handler fasthttp.RequestHandler, opts ...Option) *Server {
	httpServer := &fasthttp.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		addr:            _defaultAddr,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Start -.
func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe(s.addr)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	return s.server.Shutdown()
}
