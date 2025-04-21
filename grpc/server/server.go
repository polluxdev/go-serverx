package server

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	_defaultAddr            = ":50051"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *grpc.Server
	listener        net.Listener
	addr            string
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(server *grpc.Server, opts ...Option) (*Server, error) {
	s := &Server{
		server:          server,
		addr:            _defaultAddr,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Apply options
	for _, opt := range opts {
		opt(s)
	}

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return nil, err
	}
	s.listener = lis

	return s, nil
}

// Start -.
func (s *Server) Start() {
	log.Printf("[gRPC] server is starting on %s...", s.addr)

	go func() {
		s.notify <- s.server.Serve(s.listener)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	// GracefulStop does not block, so we simulate graceful shutdown delay
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	stopped := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		s.server.Stop()
	case <-stopped:
	}

	return nil
}
