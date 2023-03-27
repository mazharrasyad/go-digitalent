package routers

import (
	"project-2/controllers"
	"project-2/database"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	database.StartDB()
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:bookID", controllers.GetBook)
	router.GET("/books", controllers.GetAllBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
