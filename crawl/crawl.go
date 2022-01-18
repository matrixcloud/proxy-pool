package crawl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/matrixcloud/proxy-pool/common"
	"github.com/matrixcloud/proxy-pool/util"
	"github.com/rs/zerolog/log"
)

type Crawler func() []common.Proxy

// Crawlers is a map
var Crawlers = map[string]Crawler{
	"kxdaili": crawlKxdaili,
	// "kuaidaili": crawlKuaidaili,
	// "data5u":    crawlData5u,
	// "xicidaili": crawlXicidaili,
}

func crawlKuaidaili() []common.Proxy {
	result := []common.Proxy{}

	for page := 1; page <= 3; page++ {
		url := fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d", page)
		log.Info().Msgf("Start to crawl %s", url)

		doc, err := util.GetPage(url)
		if err != nil {
			log.Err(err)
			continue
		}

		doc.Find("#list > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
			ip := strings.TrimSpace(s.Find("td[data-title='IP']").Text())
			iport, _ := strconv.Atoi(strings.TrimSpace(s.Find("td[data-title='PORT']").Text()))
			port := uint16(iport)
			schema := strings.ToLower(strings.TrimSpace(s.Find("td[data-title='类型']").Text()))

			result = append(result, *common.NewProxy(schema, ip, port))
		})
	}

	return result
}

func crawlKxdaili() []common.Proxy {
	result := []common.Proxy{}
	url := "http://www.kxdaili.com/dailiip.html"
	log.Debug().Msgf("Start to crawl %s", url)

	doc, err := util.GetPage(url)
	if err != nil {
		log.Err(err)
		return nil
	}

	doc.Find(".hot-product-content > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		ip := strings.TrimSpace(s.Find(":nth-child(1)").Text())
		port, _ := strconv.Atoi(strings.TrimSpace(s.Find(":nth-child(2)").Text()))
		schema := strings.ToLower(strings.TrimSpace(s.Find(":nth-child(4)").Text()))

		result = append(result, *common.NewProxy(schema, ip, uint16(port)))
	})

	return result
}

func crawlData5u() []common.Proxy {
	result := []common.Proxy{}
	url := "http://www.data5u.com/"
	log.Printf("Start to crawl %s", url)

	doc, err := util.GetPage(url)

	if err != nil {
		log.Err(err)
		return result
	}

	doc.Find("body > div:nth-child(16) > ul > li:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		ip := strings.TrimSpace(s.Find("span:nth-child(1) > li").Text())
		port, _ := strconv.Atoi(strings.TrimSpace(s.Find("span:nth-child(2) > li").Text()))
		schema := strings.TrimSpace(s.Find("span:nth-child(4) > li").Text())

		result = append(result, *common.NewProxy(schema, ip, uint16(port)))
	})

	return result
}

func crawlXicidaili() []common.Proxy {
	result := []common.Proxy{}

	for page := 1; page <= 3; page++ {
		url := fmt.Sprintf("https://www.xicidaili.com/nn/%d", page)
		log.Printf("Start to crawl %s", url)

		doc, err := util.GetPage(url)

		if err != nil {
			log.Err(err)
			return result
		}

		doc.Find("#ip_list > tbody").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}

			ip := strings.TrimSpace(s.Find("td:nth-child(2)").Text())
			port, _ := strconv.Atoi(strings.TrimSpace(s.Find("td:nth-child(3)").Text()))
			schema := strings.TrimSpace(s.Find("td:nth-child(6)").Text())

			result = append(result, *common.NewProxy(schema, ip, uint16(port)))
		})
	}

	return result
}
