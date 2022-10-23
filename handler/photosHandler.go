package handler

import (
	"MyGram/config"
	"MyGram/entity"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostPhotoHandler(ctx *gin.Context) {
	db := config.GetDB()
	photo := entity.Photo{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	UserID := userData["id"].(string)

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Bad Request: " + err.Error(),
		})
	}

	photo.UserID = UserID
	photo.ID = uuid.New().String()

	err := db.Create(&photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Bad Request: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"statusCode": http.StatusCreated,
		"message":    "Success: photos uploaded",
		"data": entity.PostPhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
		},
	})
}
