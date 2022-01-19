package core

import (
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
)

type Validater struct {
	pool     *Pool
	interval int
}

func NewValidater(pool *Pool) *Validater {
	return &Validater{
		pool:     pool,
		interval: 1,
	}
}

func (vr *Validater) Start() {
	go func() {
		for {
			time.Sleep(time.Duration(vr.interval) * time.Second)
			validate(vr.pool)
		}
	}()

	log.Info().Msg("Validater is started")
}

const IP_CHECKER_API = "http://api.ipify.org/?format=json"
const IP_CHECKER_API_SSL = "https://api.ipify.org/?format=json"

func validate(pool *Pool) {
	for _, proxy := range pool.proxies {
		log.Debug().Msgf("Start to test %s.\n", proxy.Addr)
		url, err := url.Parse(proxy.Addr)

		if err != nil {
			log.Error().Msgf("Failed to parse " + proxy.Addr)
			proxy.Validated = false
			pool.Update(proxy)
		}

		tr := &http.Transport{
			Proxy: http.ProxyURL(url),
		}
		client := &http.Client{Transport: tr, Timeout: 1 * time.Second}
		r, err := client.Get(IP_CHECKER_API_SSL)

		if err != nil {
			proxy.Validated = false
			pool.Update(proxy)
		}

		if r.StatusCode == 200 {
			proxy.Validated = true
			pool.Update(proxy)
		}

		defer r.Body.Close()

		proxy.Validated = true
		pool.Update(proxy)
	}
}
