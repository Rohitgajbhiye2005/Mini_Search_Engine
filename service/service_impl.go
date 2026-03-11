package service

import (
	"mini_search_engine/crawler"
	"mini_search_engine/model"
	"mini_search_engine/repository"
)

type pageServiceImpl struct {
    repo repository.PageRepository
}

func NewPageService(repo repository.PageRepository) PageService {
    return &pageServiceImpl{repo: repo}
}

// CrawlAndStore now accepts seed URLs and maxPages
func (s *pageServiceImpl) CrawlAndStore(seedUrls []string, maxPages int) error {
    // The crawler now handles inserting pages directly into the repo
    crawler.Crawl(seedUrls, maxPages, s.repo)
    return nil
}

// service/page_service.go
func (s *pageServiceImpl) Search(query string, limit int) ([]model.SearchResult, error) {
    return s.repo.SearchPages(query, limit)
}