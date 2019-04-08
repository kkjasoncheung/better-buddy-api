package main

import (
	"flag"

	"github.com/kkjasoncheung/better-buddy-api/db"
	"github.com/kkjasoncheung/better-buddy-api/migrations"
	"github.com/kkjasoncheung/better-buddy-api/server"
)

func main() {
	flag.Parse()
	db := db.Init()
	migrations.MigrateSchema(db)
	server.Init()
}
