package main

import (
	"github.com/seigaalghi/gorm-crud/controllers"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/gorm-crud/models"
)

func main() {
	server := gin.Default()
	db := models.Database()
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.News{})

	server.Use(func(context *gin.Context) {
		context.Set("db", db)
	})

	server.GET("/api/v1/authors", controllers.GetAuthors)
	server.POST("/api/v1/authors", controllers.CreateAuthor)
	server.POST("/api/v1/authors/login", controllers.Login)
	server.PUT("/api/v1/authors/:id", controllers.EditAuthor)
	server.DELETE("/api/v1/authors/:id", controllers.DeleteAuthor)

	server.GET("/api/v1/newses", controllers.GetNewses)
	server.POST("/api/v1/newses", controllers.CreateNews)

	server.Run(":5000")
}
