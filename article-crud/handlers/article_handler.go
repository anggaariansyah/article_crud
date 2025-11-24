package handlers

import (
	"article-crud/models"
	"article-crud/repository"
	"article-crud/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	Repo *repository.ArticleRepository
}

func NewArticleHandler(repo *repository.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{Repo: repo}
}

func (h *ArticleHandler) GetAll(c *gin.Context) {
	articles, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) Create(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	article := models.Article{
		Title:   title,
		Content: content,
	}

	if err := h.Repo.Create(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["photos"]
	uploadedPaths, err := utils.UploadFiles(files, "uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, path := range uploadedPaths {
		photo := models.ArticlePhoto{
			ArticleID: article.ID,
			URL:       path,
		}
		h.Repo.DB.Create(&photo)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created successfully"})
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	if title != "" {
		article.Title = title
	}
	if content != "" {
		article.Content = content
	}

	if err := h.Repo.Update(article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Handle foto
	form, _ := c.MultipartForm()
	if form != nil {
		files := form.File["photos"]
		if len(files) > 0 {
			// 1. Ambil semua foto lama
			var oldPhotos []models.ArticlePhoto
			h.Repo.DB.Where("article_id = ?", article.ID).Find(&oldPhotos)

			// 2. Hapus file lama dari folder uploads/
			for _, p := range oldPhotos {
				os.Remove(p.URL)
			}

			// 3. Hapus record lama di database
			h.Repo.DB.Where("article_id = ?", article.ID).Delete(&models.ArticlePhoto{})

			// 4. Upload foto baru
			uploadedPaths, err := utils.UploadFiles(files, "uploads")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			for _, path := range uploadedPaths {
				photo := models.ArticlePhoto{
					ArticleID: article.ID,
					URL:       path,
				}
				h.Repo.DB.Create(&photo)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
