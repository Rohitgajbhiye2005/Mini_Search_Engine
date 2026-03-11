package api

import (
	"mini_search_engine/service"
	"net/http"
	"strconv"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(svc service.PageService){
	r:=gin.Default()
	r.Use(cors.Default())

	r.GET("/search",func(ctx *gin.Context) {
		query:=ctx.Query("q")
		limister:=ctx.Query("limit")

		limit,err:=strconv.Atoi(limister)
		if err!=nil{
			limit=10
		}

		results,err:=svc.Search(query,limit)

		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK,results)

	})

	r.POST("/crawl", func(c *gin.Context) {
		var req struct {
			Seeds    []string `json:"seeds"`
			MaxPages int      `json:"max_pages"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		go svc.CrawlAndStore(req.Seeds, req.MaxPages) // run in background
		c.JSON(http.StatusOK, gin.H{"message": "Crawling started"})
	})

	// Run server on port 8080
	r.Run(":8080")

}

