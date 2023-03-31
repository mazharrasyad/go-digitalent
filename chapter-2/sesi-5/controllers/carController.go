package controllers

import (
	"chapter-2/sesi-5/database"
	"chapter-2/sesi-5/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	GetAllCars godoc
//
// @Summary Get all cars
// @Description Get details of all cars data
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(ctx *gin.Context) {
	db := database.New()

	cars := []models.Car{}
	err := db.Find(&cars).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, cars)
}

// GetCarByID godoc
// @Summary Get car by car id
// @Description Get details of specific car
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetCarByID(ctx *gin.Context) {
	db := database.New()

	car := models.Car{}
	err := db.First(&car, "id=?", ctx.Param("id")).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, car)
}

// CreateCar "chapter-2/sesi-5/controllers"doc
// @Summary Create car
// @Description Create new car
// @Tags json
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "Create car"
// @Success 201 {object} models.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	db := database.New()
	request := models.Car{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	car := models.Car{
		Model: request.Model,
		Brand: request.Brand,
	}
	db.Create(&car)

	ctx.JSON(http.StatusCreated, car)
}

// UpdateCar godoc
// @Summary Update car
// @Description Update car data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [patch]
func UpdateCar(ctx *gin.Context) {
	db := database.New()
	car := models.Car{}
	request := models.Car{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	err := db.First(&car, "id=?", ctx.Param("id")).Error
	if err != nil {
		panic(err)
	}

	car = request
	db.Save(&car)

	ctx.JSON(http.StatusOK, car)
}

// DeleteCar godoc
// @Summary Delete car
// @Description Delete car data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200
// @Router /cars/{Id} [delete]
func DeleteCar(ctx *gin.Context) {
	db := database.New()
	car := models.Car{}

	err := db.First(&car, "id=?", ctx.Param("id")).Error
	if err != nil {
		panic(err)
	}

	db.Delete(&car)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success to delete car",
	})
}
