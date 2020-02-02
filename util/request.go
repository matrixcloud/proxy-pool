package util

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetPage returns
func GetPage(url string) (*goquery.Document, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", RandomUA("all"))

	r, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		newError := fmt.Errorf("Status code error: %d %s", r.StatusCode, r.Status)
		return nil, newError
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Failed to parse html document")
	}

	return doc, nil
}
