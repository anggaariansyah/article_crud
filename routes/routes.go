package routes

import (
	"article-crud/handlers"
	"article-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, articleHandler *handlers.ArticleHandler) {
	article := r.Group("/articles")
	article.Use(middlewares.BasicAuth())
	{
		article.GET("", articleHandler.GetAll)
		article.GET("/:id", articleHandler.GetByID)
		article.POST("", articleHandler.Create)
		article.PUT("/:id", articleHandler.Update)
		article.DELETE("/:id", articleHandler.Delete)
	}
}
