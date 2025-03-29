package client

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	_defaultTarget      = "localhost:50051"
	_defaultDialTimeout = 5 * time.Second
	_defaultCallTimeout = 3 * time.Second
)

// Client -.
type Client struct {
	conn        *grpc.ClientConn
	target      string
	dialTimeout time.Duration
	callTimeout time.Duration
}

// New -.
func New(opts ...Option) (*Client, error) {
	c := &Client{
		target:      _defaultTarget,
		dialTimeout: _defaultDialTimeout,
		callTimeout: _defaultCallTimeout,
	}

	// Apply functional options
	for _, opt := range opts {
		opt(c)
	}

	conn, err := grpc.NewClient(c.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c.conn = conn

	log.Printf("[gRPC] connected to %s", c.target)

	return c, nil
}

// Conn -.
func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}

// Close -.
func (c *Client) Close() error {
	return c.conn.Close()
}
