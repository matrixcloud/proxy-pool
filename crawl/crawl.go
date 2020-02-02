package crawl

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/matrixcloud/proxy-pool/util"
)

// AddressProvider is a map
var AddressProvider = map[string]func() []string{
	"kxdaili":   crawlKxdaili,
	"data5u":    crawlData5u,
	"xicidaili": crawlXicidaili,
	"kuaidaili": crawlKuaidaili,
}

func crawlKuaidaili() []string {
	result := []string{}

	for page := 1; page <= 3; page++ {
		url := fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d", page)
		log.Printf("Start to crawl %s", url)

		doc, err := util.GetPage(url)
		if err != nil {
			log.Println(err)
			continue
		}

		doc.Find("#list > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
			ip := strings.TrimSpace(s.Find("td[data-title='IP']").Text())
			port := strings.TrimSpace(s.Find("td[data-title='PORT']").Text())
			schema := strings.TrimSpace(s.Find("td[data-title='类型']").Text())

			result = append(result, fmt.Sprintf("%s://%s:%s", strings.ToLower(schema), ip, port))
		})

		time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)
	}

	return result
}

func crawlKxdaili() []string {
	result := []string{}
	url := "http://www.kxdaili.com/dailiip.html"
	log.Printf("Start to crawl %s", url)

	doc, err := util.GetPage(url)
	if err != nil {
		log.Println(err)
		return result
	}

	doc.Find(".hot-product-content > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		ip := strings.TrimSpace(s.Find(":nth-child(1)").Text())
		port := strings.TrimSpace(s.Find(":nth-child(2)").Text())
		schema := strings.TrimSpace(s.Find(":nth-child(4)").Text())

		result = append(result, fmt.Sprintf("%s://%s:%s", strings.ToLower(schema), ip, port))
	})

	return result
}

func crawlData5u() []string {
	result := []string{}
	url := "http://www.data5u.com/"
	log.Printf("Start to crawl %s", url)

	doc, err := util.GetPage(url)

	if err != nil {
		log.Println(err)
		return result
	}

	doc.Find("body > div:nth-child(16) > ul > li:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		ip := strings.TrimSpace(s.Find("span:nth-child(1) > li").Text())
		port := strings.TrimSpace(s.Find("span:nth-child(2) > li").Text())
		schema := strings.TrimSpace(s.Find("span:nth-child(4) > li").Text())

		result = append(result, fmt.Sprintf("%s://%s:%s", strings.ToLower(schema), ip, port))
	})

	return result
}

func crawlXicidaili() []string {
	result := []string{}

	for page := 1; page <= 3; page++ {
		url := fmt.Sprintf("https://www.xicidaili.com/nn/%d", page)
		log.Printf("Start to crawl %s", url)

		doc, err := util.GetPage(url)

		if err != nil {
			log.Println(err)
			return result
		}

		doc.Find("#ip_list > tbody").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}

			ip := strings.TrimSpace(s.Find("td:nth-child(2)").Text())
			port := strings.TrimSpace(s.Find("td:nth-child(3)").Text())
			schema := strings.TrimSpace(s.Find("td:nth-child(6)").Text())

			result = append(result, fmt.Sprintf("%s://%s:%s", strings.ToLower(schema), ip, port))
		})

		time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)
	}

	return result
}
