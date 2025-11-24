package main

import (
	"article-crud/config"
	"article-crud/handlers"
	"article-crud/migrations"
	"article-crud/repository"
	"article-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	migrations.Migrate(db)

	repo := repository.NewArticleRepository(db)
	handler := handlers.NewArticleHandler(repo)

	r := gin.Default()
	r.Static("/uploads", "./uploads")
	routes.SetupRoutes(r, handler)

	r.Run(":8080")
}
