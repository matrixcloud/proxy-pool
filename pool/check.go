package pool

import (
	"log"
	"time"

	"github.com/matrixcloud/proxy-pool/crawl"
)

func (p *Pool) get() {
	for _, crawl := range crawl.AddressProvider {
		addrs := crawl()
		p.Test(addrs)

		if p.isOverThreshold() {
			break
		}
	}
}

func (p *Pool) check() {
	log.Println("Pool Checker started.")

	for {
		if p.conn.Length() < int64(p.minThreshold) {
			p.get()
			time.Sleep(time.Duration(p.checkInterval) * time.Second)
		}
	}
}
