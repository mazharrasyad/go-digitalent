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

// CreateSocialMedia "chapter-3/challange-3/controllers"doc
// @Summary Create socialMedia
// @Description Create new socialMedia
// @Tags json
// @Accept json
// @Produce json
// @Param models.SocialMedia body models.SocialMedia true "Create socialMedia"
// @Success 201 {object} models.SocialMedia
// @Router /social-medias [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// ReadSocialMedia godoc
// @Summary Get socialMedia by socialMedia id
// @Description Get details of specific socialMedia
// @Tags json
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the socialMedia"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{socialMediaId} [get]
func ReadSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	SocialMediaID := c.Param("socialMediaId")

	err := db.Debug().Where("id = ?", SocialMediaID).First(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "Social Media not found",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update socialMedia
// @Description Update socialMedia data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the socialMedia"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{Id} [patch]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete socialMedia
// @Description Delete socialMedia data
// @Tags json
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the socialMedia"
// @Success 200
// @Router /social-medias/{socialMediaId} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	err := db.Model(&SocialMedia).Where("id = ? AND user_id = ?", socialMediaId, userID).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social Media has been deleted"})
}

// ReadAllSocialMedias godoc
// @Summary Get all socialMedias
// @Description Get details of all socialMedias data
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /social-medias [get]
func ReadAllSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var SocialMedia []models.SocialMedia

	err := db.Debug().Where("user_id = ?", userID).Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "No Social Media found for this user",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	c.JSON(http.StatusOK, SocialMedia)
}
