package crawler

import (
	"fmt"
	"mini_search_engine/model"
	"mini_search_engine/repository"
	"sync"
)

func Crawl(seedUrls []string, maxPages int,repo repository.PageRepository) {
	urlQueue := make(chan string, 1000)
	// queue := seedUrls
	var mu sync.Mutex
	visited := make(map[string]bool)

	count := 0
	workers := 10

	var wg sync.WaitGroup

	for _, url := range seedUrls {
		normlized:=NormalizeURL(url)
		if normlized==""{
			continue
		}
		visited[normlized]=true
		urlQueue<-normlized
	}

	// start worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(
			urlQueue,
			visited,
			&mu,
			seedUrls[0],
			maxPages,
			&count,
			&wg,
			repo,
		)
	}
	// i have to implemenet the waitgroup here

	wg.Wait()

	//close(urlQueue)

	fmt.Println("Crawling Finished.")


	//select {}

	// seen:=make(map[string]bool)

	// for _,url:=range seedUrls{
	// 	frontier=append(frontier, url)
	// }

	// count := 0
	// linkLimitPerPage:=50

	// for len(frontier) > 0 && count < maxPages {
	// 	url := frontier[0]
	// 	frontier = frontier[1:]

	// 	if visited[url] {
	// 		continue
	// 	}

	// 	visited[url] = true
	// 	fmt.Println("Crawling:", url)

	// 	html, err := FetchPage(url)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	links := ExtractLinks(url, html)
	// 	added:=0

	// 	for _,link:= range links{
	// 		// only crawl same domain
	// 		if !SameDomain(seedUrls[0],link){
	// 			continue
	// 		}

	// 		if visited[link] || contains(frontier,link){
	// 			continue
	// 		}

	// 		frontier=append(frontier, link)
	// 		//seen[link]=true

	// 		added++

	// 		if added>=linkLimitPerPage {
	// 			break
	// 		}

	// 	}

	// 	// queue = append(queue, links...)

	// 	count++
	// 	fmt.Printf("Page %d: %s\n", count, url)
	// 	fmt.Println("Queue size:", len(frontier))

	//}
}

// func contains(slice []string, item string) bool {
// 	for _, s := range slice {
// 		if s == item {
// 			return true
// 		}
// 	}
// 	return false
// }

func worker(
	urlQueue chan string,
	visited map[string]bool,
	mu *sync.Mutex,
	seed string,
	maxPages int,
	count *int,
	wg *sync.WaitGroup,
	repo repository.PageRepository,
) {
	//linkLimitPerPage := 500
	defer wg.Done()

	for {
		url := <-urlQueue
		// if !ok{
		// 	return
		// }

		mu.Lock()
		if *count >= maxPages {
			mu.Unlock()
			return	
		}

		// visited[url] = true
		// *count++
		currentCount := *count+1
		*count=currentCount
		mu.Unlock()

		fmt.Printf("Page %d: %s\n", currentCount, url)

		html, err := FetchPage(url)
		if err != nil {
			continue
		}
		title := ExtractTitle(html)
		text := ExtractText(html)
		page:=&model.Page{
			URL: url,
			Title: title,
			Content: text,
		}
		if err:=repo.InsertPage(page);err!=nil{
			fmt.Printf("Failed to insert page %s: %v\n",url,err)
		}
		links := ExtractLinks(url, html)

		//fmt.Println("Queue size:", len(urlQueue))
		//added := 0

		for _, link := range links {
			if !SameDomain(seed, link) {
				continue
			}
			link = NormalizeURL(link)
			if link == "" {
				continue
			}

			if !IsHTMLPage(link) {
				continue
			}

			mu.Lock()
			if visited[link] {
				mu.Unlock()
				continue
			}

			if *count>=maxPages{
				mu.Unlock()
				return
			} 
			visited[link]=true
			mu.Unlock()


			select{	
			case urlQueue <- link:
			default:
			}
			//added++

			// if added >= linkLimitPerPage {
			// 	break
			// }
		}
		fmt.Println("Queue size:", len(urlQueue))
	}
}

// func worker(
// 	urlQueue chan string,
// 	visited map[string]bool,
// 	mu *sync.Mutex,
// 	seed string,
// 	maxPages int,
// 	count *int,
// ) {

// 	linkLimitPerPage := 50

// 	for {
// 		url := <-urlQueue

// 		mu.Lock()
// 		if *count >= maxPages {
// 			mu.Unlock()
// 			return
// 		}

// 		if visited[url] {
// 			mu.Unlock()
// 			continue
// 		}

// 		visited[url] = true
// 		*count++
// 		currentCount := *count
// 		mu.Unlock()

// 		fmt.Printf("Page %d: %s\n", currentCount, url)

// 		html, err := FetchPage(url)
// 		if err != nil {
// 			continue
// 		}

// 		links := ExtractLinks(url, html)

// 		added := 0

// 		for _, link := range links {

// 			if !SameDomain(seed, link) {
// 				continue
// 			}

// 			mu.Lock()
// 			if visited[link] {
// 				mu.Unlock()
// 				continue
// 			}

// 			visited[link] = true
// 			mu.Unlock()

// 			urlQueue <- link

// 			added++

// 			if added >= linkLimitPerPage {
// 				break
// 			}
// 		}

// 		fmt.Println("Queue size:", len(urlQueue))
// 	}
// }
