package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kkjasoncheung/better-buddy-api/models"
)

// MigrateSchema migrates any schemas at initialization.
func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.Model(models.User{}).AddIndex("idx_first_name", "first_name")
	db.Model(models.User{}).AddIndex("idx_last_name", "last_name")

	db.AutoMigrate(models.Companion{})
	db.Model(models.User{}).AddIndex("idx_name", "name")
}
