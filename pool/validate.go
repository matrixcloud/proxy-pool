package pool

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func (p *Pool) validate() {
	log.Println("Pool validator started.")

	for {
		time.Sleep(time.Duration(p.validateInterval) * time.Second)
		count := int(p.conn.Length() / 2)

		// Wait to get new raw proxies
		if count == 0 {
			time.Sleep(time.Duration(p.validateInterval) * time.Second)
			continue
		}

		proxies := p.conn.Get(int64(count))
		p.Test(proxies)
	}
}

const testAPI = "http://www.baidu.com"

// Test checks raw proxies. If raw proxy is available, put it into queue
func (p *Pool) Test(proxies []string) {
	for _, proxy := range proxies {
		if p.testOne(proxy) {
			log.Printf("%s checked.\n", proxy)
			p.conn.Push(proxy)
		} else {
			log.Printf("Invalid %s.\n", proxy)
		}
	}
}

func (p *Pool) testOne(proxy string) bool {
	log.Printf("Start to test %s.\n", proxy)
	url, err := url.Parse(proxy)

	if err != nil {
		log.Println("Failed to parse " + proxy)
		return false
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	client := &http.Client{Transport: tr, Timeout: 1 * time.Second}
	r, err := client.Get(testAPI)

	if err != nil {
		return false
	}

	if r.StatusCode == 200 {
		return true
	}

	defer r.Body.Close()

	return false
}
