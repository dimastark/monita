package webpage

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/jaytaylor/html2text"
)

// WebPage provides JQuery-like methods
type WebPage struct {
	*goquery.Document
}

// Load the WebPage by url
func Load(url string) (*WebPage, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	d, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, err
	}

	return &WebPage{d}, nil
}

// Select gets text matched selector or empty string
func (w *WebPage) Select(selector string) string {
	html, err := w.Find(selector).First().Html()

	if err != nil {
		return ""
	}

	text, err := html2text.FromString(html, html2text.Options{PrettyTables: true})

	if err != nil {
		return ""
	}

	return text
}
