package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// Init initializes the database with proper config and credentials.
func Init() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	dbUsername := os.Getenv("MYSQL_USERNAME")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbCredentails := dbUsername + ":" + dbPassword + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open("mysql", dbCredentails)
	if err != nil {
		log.Fatal("Error connecting to database.", err)
	}
	return db
}

// GetDb returns the database connection.
func GetDb() *gorm.DB {
	return db
}
