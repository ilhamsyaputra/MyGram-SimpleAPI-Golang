package router

import (
	"MyGram/handler"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handler.UserRegisterHandler)
		userRouter.POST("/login", handler.UserLoginHandler)
	}

	return r
}
