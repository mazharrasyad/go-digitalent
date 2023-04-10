package controllers

import (
	"chapter-3/project-3/database"
	"chapter-3/project-3/helpers"
	"chapter-3/project-3/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto "chapter-3/challange-3/controllers"doc
// @Summary Create photo
// @Description Create new photo
// @Tags json
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "Create photo"
// @Success 201 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// ReadPhoto godoc
// @Summary Get photo by photo id
// @Description Get details of specific photo
// @Tags json
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [get]
func ReadPhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	PhotoID := c.Param("photoId")

	err := db.Debug().Where("id = ?", PhotoID).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "Photo not found",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	c.JSON(http.StatusOK, Photo)
}

// UpdatePhoto godoc
// @Summary Update photo
// @Description Update photo data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{Id} [patch]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo data
// @Tags json
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Success 200
// @Router /photos/{photoId} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	err := db.Model(&Photo).Where("id = ? AND user_id = ?", photoId, userID).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo has been deleted"})
}

// ReadAllPhotos godoc
// @Summary Get all photos
// @Description Get details of all photos data
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos [get]
func ReadAllPhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var Photo []models.Photo

	err := db.Debug().Where("user_id = ?", userID).Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "No Photo found for this user",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	c.JSON(http.StatusOK, Photo)
}
