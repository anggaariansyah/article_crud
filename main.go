package main

import (
	"article-crud/config"
	"article-crud/handlers"
	"article-crud/migrations"
	"article-crud/repository"
	"article-crud/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	migrations.Migrate(db)

	repo := repository.NewArticleRepository(db)
	handler := handlers.NewArticleHandler(repo)

	r := gin.Default()

	// Static folder untuk foto
	r.Static("/uploads", "./uploads")

	// Tambahkan CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup route artikel + middleware BasicAuth
	routes.SetupRoutes(r, handler)

	r.Run(":8080")
}
