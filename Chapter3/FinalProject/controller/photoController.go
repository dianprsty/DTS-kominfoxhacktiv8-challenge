package controller

import (
	"net/http"
	"time"

	"github.com/dianprsty/DTS-kominfoxhacktiv8-challenge/tree/main/Chapter3/FinalProject/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoInput struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}

// GetAllPhoto godoc
// @Summary Get all Photo.
// @Description Get a list of Photo.
// @Tags Photo
// @Produce json
// @Success 200 {object} []models.Photo
// @Router /photos [get]
func GetAllPhoto(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var Photos []models.Photo
	db.Find(&Photos)

	c.JSON(http.StatusOK, gin.H{"data": Photos})
}

// CreatePhoto godoc
// @Summary Create New Photo.
// @Description Creating a new Photo.
// @Tags Photo
// @Param Body body PhotoInput true "the body to create a new Photo"
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	// Validate input
	var input PhotoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo := models.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
		UserId:   input.UserId,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&photo)

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

// GetPhotoById godoc
// @Summary Get Photo.
// @Description Get an Photo by id.
// @Tags Photo
// @Produce json
// @Param id path string true "Photo id"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [get]
func GetPhotoById(c *gin.Context) { // Get model if exist
	var Photo models.Photo

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Photo})
}

// UpdatePhoto godoc
// @Summary Update Photo.
// @Description Update Photo by id.
// @Tags Photo
// @Produce json
// @Param id path string true "Photo id"
// @Param Body body PhotoInput true "the body to update age Photo category"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [put]
func UpdatePhoto(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Photo models.Photo
	if err := db.Where("id = ?", c.Param("id")).First(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input PhotoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Photo
	updatedInput.Title = input.Title
	updatedInput.Caption = input.Caption
	updatedInput.PhotoUrl = input.PhotoUrl
	updatedInput.UpdatedAt = time.Now()

	db.Model(&Photo).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Photo})
}

// DeletePhoto godoc
// @Summary Delete one Photo.
// @Description Delete a Photo by id.
// @Tags Photo
// @Produce json
// @Param id path string true "Photo id"
// @Success 200 {object} map[string]boolean
// @Router /photos/{id} [delete]
func DeletePhoto(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var Photo models.Photo
	if err := db.Where("id = ?", c.Param("id")).First(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&Photo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
