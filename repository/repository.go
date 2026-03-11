package repository

import (
	"mini_search_engine/model"
)

type PageRepository interface{
	InsertPage(page *model.Page) error
	SearchPages(query string,limit int)([]model.SearchResult,error)
}

