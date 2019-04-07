package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kkjasoncheung/better-buddy-api/db"
	"github.com/kkjasoncheung/better-buddy-api/migrations"
	"github.com/kkjasoncheung/better-buddy-api/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	db := db.Init()
	migrations.MigrateSchema(db)
	server.Init()
}
