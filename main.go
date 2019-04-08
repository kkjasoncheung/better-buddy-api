package main

import (
	"flag"

	"github.com/kkjasoncheung/better-buddy-api/db"
	"github.com/kkjasoncheung/better-buddy-api/migrations"
	"github.com/kkjasoncheung/better-buddy-api/server"
)

func main() {
	flag.Parse()
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	db := db.Init()
	migrations.MigrateSchema(db)
	server.Init()
}
