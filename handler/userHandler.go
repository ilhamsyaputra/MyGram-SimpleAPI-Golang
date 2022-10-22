package handler

import (
	"MyGram/config"
	"MyGram/entity"
	"MyGram/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EmailRegistered(userRegister *entity.User) bool {
	db := config.GetDB()

	rowsAffected := db.First(&userRegister, "email = ?", userRegister.Email).RowsAffected

	if rowsAffected == 1 {
		return true
	}
	return false
}

func UsernameTaken(userRegister *entity.User) bool {
	db := config.GetDB()

	rowsAffected := db.First(&userRegister, "username = ?", userRegister.Username).RowsAffected

	if rowsAffected == 1 {
		return true
	}
	return false
}

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

	if EmailRegistered(&userRegister) {
		ctx.JSON(http.StatusConflict, gin.H{
			"statusCode": http.StatusConflict,
			"message":    "Conflict: Email already registered",
		})
		return
	}

	if UsernameTaken(&userRegister) {
		ctx.JSON(http.StatusConflict, gin.H{
			"statusCode": http.StatusConflict,
			"message":    "Conflict: Username already taken",
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
