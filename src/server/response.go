package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseApiF(bodyFunc func(context *gin.Context) (any, error)) func(context *gin.Context) {
	return func(context *gin.Context) {
		ResponseApi(context, bodyFunc)
	}
}

func ResponseApi(context *gin.Context, bodyFunc func(context *gin.Context) (any, error)) {
	data, err := bodyFunc(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    "500",
			"success": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"success": true,
		"data":    data,
	})
}

func ResponseApiError(context *gin.Context, message string) {
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"code":    "500",
		"success": false,
	})
}
