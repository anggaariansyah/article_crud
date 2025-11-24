package migrations

import (
	"article-crud/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Article{}, &models.ArticlePhoto{})
}
