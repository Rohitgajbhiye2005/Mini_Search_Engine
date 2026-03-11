package main

import (
	"log"
	"mini_search_engine/api"
	"mini_search_engine/config"
	//crawler "mini_search_engine/crawler"
	"mini_search_engine/db"
	"mini_search_engine/repository"
	"mini_search_engine/service"
)

func main() {

	cfg:=config.Load()
	dbConn:=db.Initializer(cfg)
	defer dbConn.Close()

	repo:=repository.NewPostgresRepository(dbConn)
	svc:=service.NewPageService(repo)
	// seeds := []string{
	// 	//"https://youtube.com",
	// 	"https://golang.org",
	// 	"https://stackoverflow.com",
	// }
	log.Println("Starting Rest Api server on :8080")
	api.StartServer(svc)
	//crawler.Crawl(seeds, 1000,repo)
}
