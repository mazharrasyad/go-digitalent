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

// CreateComment "chapter-3/challange-3/controllers"doc
// @Summary Create comment
// @Description Create new comment
// @Tags json
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "Create comment"
// @Success 201 {object} models.Comment
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// ReadComment godoc
// @Summary Get comment by comment id
// @Description Get details of specific comment
// @Tags json
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [get]
func ReadComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	CommentID := c.Param("commentId")

	err := db.Debug().Where("id = ?", CommentID).First(&Comment).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "Comment not found",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	c.JSON(http.StatusOK, Comment)
}

// UpdateComment godoc
// @Summary Update comment
// @Description Update comment data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the comment"
// @Success 200 {object} models.Comment
// @Router /comments/{Id} [patch]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Debug().Model(&Comment).Where("id = ?", commentId).Updates(map[string]interface{}{
		"message": Comment.Message,
	}).Preload("Photo").First(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment data
// @Tags json
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Success 200
// @Router /comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	err := db.Model(&Comment).Where("id = ? AND user_id = ?", commentId, userID).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment has been deleted"})
}

// ReadAllComments godoc
// @Summary Get all comments
// @Description Get details of all comments data
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments [get]
func ReadAllComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var Comment []models.Comment

	err := db.Debug().Where("user_id = ?", userID).Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "No Comment found for this user",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	c.JSON(http.StatusOK, Comment)
}
