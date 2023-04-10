package router

import (
	"chapter-3/project-3/controllers"
	_ "chapter-3/project-3/docs"
	"chapter-3/project-3/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description Service to manage data
// @termsOfService https://google.com
// @contact.name API Support
// @contact.email muhazharrasyad@gmail.com
// @lisence.name Apache 2.0
// @lisence.url https://google.com
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/:photoId", middlewares.PhotoAuthorization(), controllers.ReadPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
		photoRouter.GET("/", middlewares.PhotoAuthorization(), controllers.ReadAllPhoto)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.ReadSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
		socialMediaRouter.GET("/", middlewares.SocialMediaAuthorization(), controllers.ReadAllSocialMedia)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorization(), controllers.ReadComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
		commentRouter.GET("/", middlewares.CommentAuthorization(), controllers.ReadAllComment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
