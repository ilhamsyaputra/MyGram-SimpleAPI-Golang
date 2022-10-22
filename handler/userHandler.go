package handler

import (
	"MyGram/config"
	"MyGram/entity"
	"MyGram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserRegisterHandler(ctx *gin.Context) {
	db := config.GetDB()
	userRegister := entity.User{}

	if err := ctx.ShouldBindJSON(&userRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Bad Request",
		})
		return
	}

	userRegister.ID = uuid.New().String()
	userRegister.Password = helper.HashPass(userRegister.Password)

	err := db.Create(&userRegister).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Something is wrong",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"statusCode": http.StatusCreated,
		"message":    "User registration success",
		"data": entity.UserRegisterResponse{
			ID:       userRegister.ID,
			Username: userRegister.Username,
			Email:    userRegister.Email,
			Age:      userRegister.Age,
		},
	})
}

func UserLoginHandler(ctx *gin.Context) {
	db := config.GetDB()
	user := entity.User{}
	password := ""

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Bad Request",
		})
		return
	}

	password = user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "Unauthorized: invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "Unauthorized: invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
