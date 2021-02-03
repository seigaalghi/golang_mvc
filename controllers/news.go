package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/gorm-crud/models"
	"gorm.io/gorm"
)

// NewsInput is for data validation
type NewsInput struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"authorId"`
}

// GetNewses is to get all news data from database
func GetNewses(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)
	var newses []models.News
	var author models.Author
	db.Model(&newses).Association("Author").Find(&author)
	context.JSON(200, gin.H{"data": newses})
}

// CreateNews is to create or post a news
func CreateNews(context *gin.Context) {
	var input NewsInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	news := models.News{Title: input.Title, Content: input.Content, AuthorID: input.AuthorID}
	db := context.MustGet("db").(*gorm.DB)
	db.Create(&news)
	context.JSON(200, gin.H{"data": news})
}
