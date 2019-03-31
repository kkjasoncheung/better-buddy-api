package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kkjasoncheung/better-buddy-api/db"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	db.Init()
	r.Run() // listen and serve on 0.0.0.0:8080
}
