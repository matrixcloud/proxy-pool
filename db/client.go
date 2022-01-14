package db

import (
	"sync"
	"time"
)

// Client provides a API to manipulate proxies
type Client struct {
	mutex   sync.RWMutex
	proxies []Proxy
	// proxy address to index
	indexer map[string]int
}

func NewClient() *Client {
	return &Client{
		proxies: make([]Proxy, 0),
		indexer: make(map[string]int),
	}
}

func (client *Client) Add(proxy Proxy) {
	addr := proxy.Addr()
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if idx, ok := client.indexer[addr]; ok {
		found := client.proxies[idx]
		found.updatedAt = time.Now().UnixMilli()
	} else {
		proxy.createdAt = time.Now().UnixMilli()
		proxy.updatedAt = time.Now().UnixMilli()
		client.proxies = append(client.proxies, proxy)
	}
}

func (client *Client) Update(proxy Proxy) {
	addr := proxy.Addr()

	if idx, ok := client.indexer[addr]; ok {
		found := client.proxies[idx]
		found.updatedAt = time.Now().UnixMilli()
	}
}

func (client *Client) Remove(proxy Proxy) {
	addr := proxy.Addr()

	if idx, ok := client.indexer[addr]; ok {
		client.proxies = append(client.proxies[:idx], client.proxies[idx+1:]...)
	}
}

func (client *Client) Get(count int) []Proxy {
	return make([]Proxy, 0)
}

func (client *Client) Length() int {
	return len(client.proxies)
}
