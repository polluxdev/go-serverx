package server

import "time"

type Option func(*Server)

func Addr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
		s.server.Addr = addr
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
