package routers

import (
	"chapter-2/sesi-5/controllers"
	_ "chapter-2/sesi-5/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @description Service to manage car data
// @termsOfService https://google.com
// @contact.name API Support
// @contact.email admin@mail.me
// @lisence.name Apache 2.0
// @lisence.url https://google.com
// @host localhost:3000
// @BasePath /
func Router() *gin.Engine {
	r := gin.Default()

	// Read All
	r.GET("/cars", controllers.GetAllCars)
	// Read
	r.GET("/cars/:id", controllers.GetCarByID)
	// Create
	r.POST("/cars", controllers.CreateCar)
	// Update
	r.PATCH("/cars/:id", controllers.UpdateCar)
	// Delete
	r.DELETE("/cars/:id", controllers.DeleteCar)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
