package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kkjasoncheung/better-buddy-api/db"
	"github.com/kkjasoncheung/better-buddy-api/migrations"
	"github.com/kkjasoncheung/better-buddy-api/models"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		user := new(models.User)
		user.CreateUser()
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Created user.",
		})
	})
	db := db.Init()
	migrations.MigrateSchema(db)
	r.Run() // listen and serve on 0.0.0.0:8080
}
