package handler

import (
	"MyGram/config"
	"MyGram/entity"
	"MyGram/helper"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
			"message":    "Bad Request: " + err.Error(),
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

	err := db.Create(&userRegister).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Bad Request: " + err.Error(),
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

func UserUpdateHandler(ctx *gin.Context) {
	db := config.GetDB()
	username := ctx.Param("username")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := entity.User{}
	UserID := userData["id"].(string)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Bad Request",
		})
		return
	}

	if EmailRegistered(&user) {
		ctx.JSON(http.StatusConflict, gin.H{
			"statusCode": http.StatusConflict,
			"message":    "Conflict: Email already registered",
		})
		return
	}

	if UsernameTaken(&user) {
		ctx.JSON(http.StatusConflict, gin.H{
			"statusCode": http.StatusConflict,
			"message":    "Conflict: Username already taken",
		})
		return
	}

	user.ID = UserID

	err := db.Model(&user).Where("username = ?", username).Updates(
		entity.User{
			Email:    user.Email,
			Username: user.Username,
		},
	).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Error: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success update user",
		"data": entity.UserUpdateResponse{
			ID:        user.ID,
			Username:  user.Email,
			Email:     user.Email,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

func DeleteUserHandler(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	UserID := userData["id"].(string)

	comments := entity.Comment{}
	socialMedia := entity.SocialMedia{}
	photos := entity.Photo{}
	user := entity.User{}

	deleteAction := db.Where("user_id = ?", UserID).Delete(&socialMedia, &comments, &photos)

	if deleteAction.Error != nil {
		fmt.Println("Error deleting user:", deleteAction.Error.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Error: " + deleteAction.Error.Error(),
		})
	} else {
		db.Where("id = ?", UserID).Delete(&user)
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "Your account has been successfully deleted",
		})
		return
	}
}
