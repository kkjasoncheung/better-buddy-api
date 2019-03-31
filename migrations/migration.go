package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kkjasoncheung/better-buddy-api/models"
)

// MigrateSchema migrates any schemas at initialization.
func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
