package database

import (
	"github.com/Smarttin/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates a comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
