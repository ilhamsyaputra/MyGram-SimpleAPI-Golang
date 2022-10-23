package router

import (
	"MyGram/handler"
	"MyGram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handler.UserRegisterHandler)
		userRouter.POST("/login", handler.UserLoginHandler)
		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/:username", middleware.Authorization(), handler.UserUpdateHandler)
		userRouter.DELETE("/", middleware.DeleteAuthorization(), handler.DeleteUserHandler)
	}

	return r
}
