package routers

import (
	"MyGram/handler"
	"MyGram/middleware"
	"MyGram/repository"
	"MyGram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := &repository.UserRepo{DB: db}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	router.POST("/users/login", userHandler.Login)
	router.POST("/users/register", userHandler.Register)

	router.Use(middleware.Authentication())

	userRouter := router.Group("/users")
	userRouter.PUT("/:userId", userHandler.Update)
	userRouter.DELETE("", userHandler.DeleteWithoutID)
	userRouter.DELETE("/:userId", userHandler.Delete)

	photoRepo := &repository.PhotoRepo{DB: db}
	photoService := &service.PhotoService{PhotoRepo: photoRepo}
	photoHandler := &handler.PhotoHandler{PhotoService: photoService}

	photoRouter := router.Group("/photos")
	photoRouter.GET("", photoHandler.Get)
	photoRouter.POST("", photoHandler.Create)
	photoRouter.Use(middleware.PhotoAuthorization(photoService))
	{
		photoRouter.PUT("/:photoId", photoHandler.Update)
		photoRouter.DELETE("/:photoId", photoHandler.Delete)
	}

	commentRepo := &repository.CommentRepo{DB: db}
	commentService := &service.CommentService{CommentRepo: commentRepo}
	commentHandler := &handler.CommentHandler{CommentService: commentService}

	commentRouter := router.Group("/comments")
	commentRouter.GET("", commentHandler.Get)
	commentRouter.POST("", commentHandler.Create)
	commentRouter.Use(middleware.CommentAuthorization(commentService))
	{
		commentRouter.PUT("/:commentId", commentHandler.Update)
		commentRouter.DELETE("/:commentId", commentHandler.Delete)
	}

	socialMediaRepo := &repository.SocialMediaRepo{DB: db}
	socialMediaService := &service.SocialMediaService{SocialMediaRepo: socialMediaRepo}
	socialMediaHandler := &handler.SocialMediaHandler{SocialMediaService: socialMediaService}

	socialMediaRouter := router.Group("/socialmedias")
	socialMediaRouter.GET("", socialMediaHandler.Get)
	socialMediaRouter.POST("", socialMediaHandler.Create)
	socialMediaRouter.Use(middleware.SocialMediaAuthorization(socialMediaService))
	{
		socialMediaRouter.PUT("/:socialMediaId", socialMediaHandler.Update)
		socialMediaRouter.DELETE("/:socialMediaId", socialMediaHandler.Delete)
	}

	return router
}
