# Mini Search Engine (Go)

A concurrent web crawler and search engine built using Go.  
The system crawls developer documentation pages, indexes their content, and allows users to search through the indexed dataset.

This project was built to understand how real search engines work internally — including crawling, indexing, and fast query retrieval.

## Features

- Concurrent web crawler using **goroutines**
- Worker pool for parallel page fetching
- URL discovery and deduplication
- Crawls and indexes **10,000+ pages**
- Content stored in **PostgreSQL**
- Fast search using **PostgreSQL Full-Text Search (tsvector + GIN index)**
- REST API built with **Gin**
- Simple frontend search interface

## Architecture

Crawler → Content Extraction → Database Storage → Search Index → API → Frontend

## Tech Stack

- Go (Golang)
- Gin Web Framework
- PostgreSQL
- Goroutines & Channels
- Worker Pool Pattern
- PostgreSQL Full-Text Search

## Running the Project

### 1. Run backend server

go run main.go

### 2. Start crawling pages

curl -X POST http://localhost:8080/crawl \
-H "Content-Type: application/json" \
-d '{
  "seeds": ["https://golang.org"],
  "max_pages": 10000
}'

### 3. Start frontend

cd frontend  
python3 -m http.server 3000

Open in browser:

http://localhost:3000

## Example Search Queries

channel  
goroutine  
concurrency  

## Repository

https://github.com/Rohitgajbhiye2005/Mini_Search_Engine
