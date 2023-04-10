package controllers

import (
	"chapter-3/project-3/database"
	"chapter-3/project-3/helpers"
	"chapter-3/project-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// UserRegister "chapter-3/challange-3/controllers"doc
// @Summary User Register
// @Description User Register
// @Tags json
// @Accept json
// @Produce json
// @Param models.User body models.User true "User Register"
// @Success 201 {object} models.User
// @Router /register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"username": User.Username,
		"email":    User.Email,
		"age":      User.Age,
	})
}

// UserLogin "chapter-3/challange-3/controllers"doc
// @Summary User Login
// @Description User Login
// @Tags json
// @Accept json
// @Produce json
// @Param models.User body models.User true "User Login"
// @Success 201 {object} models.User
// @Router /register [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Username, User.Email, User.Age)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
