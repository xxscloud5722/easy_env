package routers

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func identityFilter(context *gin.Context) {
	var ip = context.ClientIP()
	accessToken := context.Query("access-token")
	if accessToken == "" {
		var authorization = context.GetHeader("Authorization")
		if authorization != "" && strings.HasPrefix(authorization, "Bearer ") {
			authorization = authorization[7:]
		}
		accessToken = authorization
	}
	if strings.HasPrefix(ip, "127.0.0.1") || accessToken == os.Getenv("AccessToken") {
		context.Next()
	} else {
		responseApiError(context, "401", "AccessToken Invalid")
		context.Abort()
	}
}
