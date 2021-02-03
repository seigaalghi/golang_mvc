package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/gorm-crud/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// HashPassword to hash a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPassword to check password
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetAuthors is to get all author data from database
func GetAuthors(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)
	var authors []models.Author
	db.Omit("password").Find(&authors)
	context.JSON(200, gin.H{"data": authors})
}

// CreateAuthor is to create or post an author
func CreateAuthor(context *gin.Context) {
	var input models.Author

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	password, err := HashPassword(input.Password)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{Name: input.Name, Email: input.Email, Password: password}
	db := context.MustGet("db").(*gorm.DB)
	db.Create(&author)
	context.JSON(200, gin.H{"data": author})
}

// Login ...
func Login(context *gin.Context) {
	var input models.LoginForm
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
	} else {
		var author models.Author
		db := context.MustGet("db").(*gorm.DB)
		db.Where("email = ?", input.Email).Find(&author)
		if password := CheckPassword(input.Password, author.Password); password {
			context.JSON(http.StatusOK, gin.H{
				"data": author,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"message": "Invalid Email or Password",
			})
		}
	}
}

// EditAuthor is to edit author information
func EditAuthor(context *gin.Context) {

	var author models.Author
	db := context.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", context.Param("id")).Find(&author).Error; err != nil {
		context.JSON(400, gin.H{"error": "No Author Found"})
		return
	}

	var input models.Author
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Model(&author).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": author})
}

// DeleteAuthor ...
func DeleteAuthor(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)
	var author models.Author

	if err := db.Where("id = ?", context.Param("id")).First(&author).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Author not found!"})
		return
	}
	db.Delete(&author)
	context.JSON(http.StatusOK, gin.H{"data": "Author deleted successfully"})
}
