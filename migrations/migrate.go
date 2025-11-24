package migrations

import (
	"article-crud/models"

	"gorm.io/gorm"
)

type ArticlePhoto struct {
	ID        uint   `gorm:"primaryKey"`
	ArticleID uint   `gorm:"not null;index"`
	URL       string `gorm:"not null"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Article{}, &models.ArticlePhoto{})
}
