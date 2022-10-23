package middleware

import (
	"MyGram/config"
	"MyGram/entity"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PutAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := config.GetDB()
		username := ctx.Param("username")

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		UserID := userData["id"]
		User := entity.User{}

		err := db.Select("id").First(&User, "username = ?", username).Error

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"statusCode": http.StatusNotFound,
				"message":    "Error: Data doesn't exist",
			})
		}

		if User.ID != UserID {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"statusCode": http.StatusUnauthorized,
				"message":    "You are not allowed to do this action",
			})
		}

		ctx.Next()
	}
}

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := config.GetDB()
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		UserID := userData["id"]
		User := entity.User{}

		err := db.Select("id").First(&User, "id = ?", UserID).Error

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"statusCode": http.StatusNotFound,
				"message":    "Error: Data doesn't exist",
			})
		}

		if User.ID != UserID {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"statusCode": http.StatusUnauthorized,
				"message":    "You are not allowed to do this action",
			})
		}

		ctx.Next()
	}
}
