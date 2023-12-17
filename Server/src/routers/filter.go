package routers

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func identityFilter(context *gin.Context) {
	var ip = context.ClientIP()
	accessToken := context.Query("access-token")
	if strings.HasPrefix(ip, "127.0.0.1") || accessToken == os.Getenv("AccessToken") {
		context.Next()
	} else {
		responseApiError(context, "AccessToken Invalid")
		context.Abort()
	}
}
