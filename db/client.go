package db

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
)

// PROXIES is the proxies queue name
const PROXIES = "proxies"

// Client provides a API to manipulate proxies queue
type Client struct {
	conn *redis.Client
}

// Options used for client config
type Options struct {
	Host string
	Port uint
	Pass string
}

// NewClient creates a redis client agent
func NewClient(options *Options) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", options.Host, options.Port),
		Password: options.Pass,
		DB:       0,
	})
	return &Client{
		conn: client,
	}
}

// Test checks redis connection
func Test(c *Client) bool {
	_, err := c.conn.Ping().Result()
	if err != nil {
		return false
	}

	return true
}

// Get gets proxy addresses
func (c *Client) Get(count int64) []string {
	proxies := c.conn.LRange(PROXIES, 0, count-1).Val()
	c.conn.LTrim(PROXIES, count, -1)
	return proxies
}

// Push places a new proxy to queue
func (c *Client) Push(proxy string) {
	c.conn.LPush(PROXIES, proxy)
}

// Pop pops a proxy from queue
func (c *Client) Pop() (string, error) {
	cmd := c.conn.RPop(PROXIES)
	if cmd.Err() != nil {
		return "", errors.New("Proxy pool is empty")
	}

	return cmd.Val(), nil
}

// Length returns queue size
func (c *Client) Length() int64 {
	return c.conn.LLen(PROXIES).Val()
}

// Clear deletes all items
func (c *Client) Clear() {
	c.conn.Del(PROXIES)
}
