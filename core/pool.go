package core

import (
	"sync"
	"time"

	"github.com/matrixcloud/proxy-pool/common"
)

// Pool used for manipulate proxies pool periodicly
type Pool struct {
	mutex   sync.RWMutex
	proxies []common.Proxy
	// proxy address to index
	indexer map[string]int
	// proxies settings
	minThreshold int
	maxThreshold int
}

// NewPool returns a proxy pool instance
func NewPool(minThreshold int, maxThreshold int) *Pool {
	return &Pool{
		proxies:      make([]common.Proxy, 0),
		indexer:      make(map[string]int),
		minThreshold: minThreshold,
		maxThreshold: maxThreshold,
	}
}

func (pool *Pool) Add(proxy common.Proxy) {
	addr := proxy.Addr
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if idx, ok := pool.indexer[addr]; ok {
		found := pool.proxies[idx]
		found.UpdatedAt = time.Now().UnixMilli()
	} else {
		proxy.CreatedAt = time.Now().UnixMilli()
		proxy.UpdatedAt = time.Now().UnixMilli()
		pool.proxies = append(pool.proxies, proxy)
	}
}

func (pool *Pool) AddAll(proxies []common.Proxy) {
	for _, proxy := range proxies {
		pool.Add(proxy)
	}
}

func (pool *Pool) Update(proxy common.Proxy) {
	addr := proxy.Addr

	if idx, ok := pool.indexer[addr]; ok {
		found := pool.proxies[idx]
		found.UpdatedAt = time.Now().UnixMilli()
	}
}

func (pool *Pool) Remove(proxy common.Proxy) {
	addr := proxy.Addr

	if idx, ok := pool.indexer[addr]; ok {
		pool.proxies = append(pool.proxies[:idx], pool.proxies[idx+1:]...)
	}
}

func (pool *Pool) Get(count int) []common.Proxy {
	return pool.proxies
}

func (pool *Pool) Length() int {
	return len(pool.proxies)
}

func (pool *Pool) IsOverThreadhold() bool {
	return len(pool.proxies) > pool.maxThreshold
}
