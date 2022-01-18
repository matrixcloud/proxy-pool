package core

import (
	"github.com/matrixcloud/proxy-pool/common"
	"github.com/matrixcloud/proxy-pool/crawl"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	pool             *Pool
	shutdown         chan struct{}
	maxThreshold     int
	minThreshold     int
	checkInterval    int
	validateInterval int
}

func NewScheduler(pool *Pool) *Scheduler {
	return &Scheduler{
		pool:     pool,
		shutdown: make(chan struct{}),
	}
}

func (shed *Scheduler) Start() {
	go shed.startFetcher()
	go shed.startValidater()
	log.Info().Msg("Scheduler is started")
}

func (shed *Scheduler) Shutdown() {
	log.Info().Msg("Scheduler is shutdown")
}

func (shed *Scheduler) startFetcher() {
	done := make(chan struct{})
	proxies := make(chan []common.Proxy)
	pool := shed.pool

	for {
		if pool.IsOverThreadhold() {
			continue
		}

		for _, crawler := range crawl.Crawlers {
			go func(crawler crawl.Crawler) {
				select {
				case proxies <- crawler():
					if pool.IsOverThreadhold() {
						close(done)
					}
					pool.AddAll(<-proxies)
				case <-done:
				}
			}(crawler)
		}
	}
}

func (shed *Scheduler) startValidater() {
	for {

	}
}

func (shed *Scheduler) validate(proxy common.Proxy) bool {
	return false
}
