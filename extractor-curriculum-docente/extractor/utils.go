package extractor

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getUrl(username string) string {
	return UrlBuscador + username
}

func getDocumentFromUrl(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil
	}

	return doc
}
