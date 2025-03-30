package client

import "time"

// Option -.
type Option func(*Client)

// Target -.
func Target(target string) Option {
	return func(c *Client) {
		c.target = target
	}
}

// DialTimeout -.
func DialTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.dialTimeout = timeout
	}
}

// CallTimeout -.
func CallTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.callTimeout = timeout
	}
}
