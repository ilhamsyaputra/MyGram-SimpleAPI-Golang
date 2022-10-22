package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ini test aja bang",
	})
}
