package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func responseApiF(bodyFunc func(context *gin.Context) (any, error)) func(context *gin.Context) {
	return func(context *gin.Context) {
		responseApi(context, bodyFunc)
	}
}

func responseApi(context *gin.Context, bodyFunc func(context *gin.Context) (any, error)) {
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

func responseApiError(context *gin.Context, message string) {
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"code":    "500",
		"success": false,
	})
}
