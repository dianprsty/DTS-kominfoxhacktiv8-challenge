package controller

import (
	"chap2-project2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookInput struct {
	Name   string `json:"name_book"`
	Author string `json:"author"`
}

// GetAllBook godoc
// @Summary Get all Book.
// @Description Get a list of Book.
// @Tags Book
// @Produce json
// @Success 200 {object} []models.Book
// @Router /books [get]
func GetAllBook(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var ratings []models.Book
	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// CreateBook godoc
// @Summary Create New Book.
// @Description Creating a new Book.
// @Tags Book
// @Param Body body BookInput true "the body to create a new Book"
// @Produce json
// @Success 200 {object} models.Book
// @Router /books [post]
func CreateBook(c *gin.Context) {
	// Validate input
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Book
	book := models.Book{Name: input.Name, Author: input.Author}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GetBookById godoc
// @Summary Get Book.
// @Description Get an Book by id.
// @Tags Book
// @Produce json
// @Param id path string true "Book id"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func GetBookById(c *gin.Context) { // Get model if exist
	var book models.Book

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary Update Book.
// @Description Update Book by id.
// @Tags Book
// @Produce json
// @Param id path string true "Book id"
// @Param Body body BookInput true "the body to update book"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Book
	updatedInput.Name = input.Name
	updatedInput.Author = input.Author
	updatedInput.UpdatedAt = time.Now()

	db.Model(&book).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete one Book.
// @Description Delete a Book by id.
// @Tags Book
// @Produce json
// @Param id path string true "Book id"
// @Success 200 {object} map[string]boolean
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
