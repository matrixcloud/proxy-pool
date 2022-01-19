package core

import (
	"github.com/matrixcloud/proxy-pool/crawl"
	"github.com/rs/zerolog/log"
)

type Fetcher struct {
	pool     *Pool
	shutdown chan struct{}
}

func NewFetcher(pool *Pool) *Fetcher {
	return &Fetcher{
		pool:     pool,
		shutdown: make(chan struct{}),
	}
}

func (fr *Fetcher) Start() {
	pool := fr.pool

	go func() {
		for {
			if pool.IsOverThreadhold() {
				continue
			}

			go fetch(pool)
		}
	}()

	log.Info().Msg("Fetcher is started")
}

func (fr *Fetcher) Shutdown() {
	log.Info().Msg("Fetcher is shutdown")
}

func fetch(pool *Pool) {
	for _, crawler := range crawl.Crawlers {
		go func(crawler crawl.Crawler) {
			pool.AddAll(crawler())
		}(crawler)
	}
}
