package crawler

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func ExtractLinks(base string,html string) []string {
	var links []string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		return links
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if !exists {
			return
		}

		if strings.HasPrefix(href, "#") {
			return
		}

		if strings.HasPrefix(href, "javascript:") {
			return
		}

		if strings.HasPrefix(href, "mailto:") {
			return
		}

		fullURL:=ResolveURL(base,href)

		links = append(links,fullURL)
	})

	return links

}


func ResolveURL(base string, href string) string {
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}

	refURL, err := url.Parse(href)
	if err != nil {
		return ""
	}

	return baseURL.ResolveReference(refURL).String()
}


func SameDomain(base string, target string) bool {

	baseURL, err1 := url.Parse(base)
	targetURL, err2 := url.Parse(target)

	if err1 != nil || err2 != nil {
		return false
	}

	return baseURL.Host == targetURL.Host
}
