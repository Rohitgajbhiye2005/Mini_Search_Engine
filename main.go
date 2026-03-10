package main

import crawler "mini_search_engine/crawler"

func main() {
	seeds := []string{
		"https://golang.org",
		"https://stackoverflow.com",
	}

	crawler.Crawl(seeds, 1000)
}
