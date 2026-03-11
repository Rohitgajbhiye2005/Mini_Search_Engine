package model

import "time"

type Page struct{
	ID int
	URL string
	Title string
	Content string
	CrawledAt time.Time
}

type SearchResult struct{
	URL string
	Title string
}