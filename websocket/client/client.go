package client

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	_defaultTarget      = "ws://localhost:8082/ws"
	_defaultDialTimeout = 5 * time.Second
	_defaultCallTimeout = 3 * time.Second
)

// Client -.
type Client struct {
	conn        *websocket.Conn
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

	conn, _, err := websocket.DefaultDialer.Dial(c.target, nil)
	if err != nil {
		return nil, err
	}

	c.conn = conn

	log.Printf("[websocket] connected to %s", c.target)

	return c, nil
}

// Conn -.
func (c *Client) Conn() *websocket.Conn {
	return c.conn
}

// Close -.
func (c *Client) Close() error {
	return c.conn.Close()
}
