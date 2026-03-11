package service

import "mini_search_engine/model"

//import"mini_search_engine/model"

type PageService interface {
    CrawlAndStore(seedUrls []string, maxPages int) error
    Search(query string,limit int) ([]model.SearchResult, error)
}