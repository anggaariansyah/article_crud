package repository

import (
	"article-crud/models"
	"fmt"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{DB: db}
}

func (r *ArticleRepository) GetAll() ([]models.Article, error) {
	var articles []models.Article
	if err := r.DB.Preload("Photos").Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticleRepository) GetByID(id uint) (*models.Article, error) {
	var article models.Article
	if err := r.DB.Preload("Photos").First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) Create(article *models.Article) error {
	var count int64
	// cek apakah ada artikel dengan judul sama
	r.DB.Model(&models.Article{}).Where("title = ?", article.Title).Count(&count)
	if count > 0 {
		return fmt.Errorf("artikel dengan judul '%s' sudah ada", article.Title)
	}
	return r.DB.Create(article).Error
}

func (r *ArticleRepository) Update(article *models.Article) error {
	return r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(article).Error
}

func (r *ArticleRepository) Delete(id uint) error {
	// hapus dulu semua photo terkait
	if err := r.DB.Where("article_id = ?", id).Delete(&models.ArticlePhoto{}).Error; err != nil {
		return err
	}
	// baru hapus artikelnya
	return r.DB.Delete(&models.Article{}, id).Error
}
