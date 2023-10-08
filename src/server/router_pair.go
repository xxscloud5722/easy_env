package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nuwa/server.v3/bean"
	"github.com/nuwa/server.v3/service"
	"os"
	"strings"
)

func (server *Server) LoadPair() {
	api := server.Group("/pair")
	api.Use(func(context *gin.Context) {
		var ip = context.Request.RemoteAddr
		accessToken := context.Query("access-token")
		if strings.HasPrefix(ip, "127.0.0.1") || accessToken == os.Getenv("AccessToken") {
			context.Next()
		} else {
			ResponseApiError(context, "access-token invalid")
			context.Abort()
		}
	})
	var pairService = service.Pair{}
	api.GET("/:key", ResponseApiF(func(context *gin.Context) (any, error) {
		var key = context.Param("key")
		return pairService.KeyValue(key)
	}))
	api.GET("/list", ResponseApiF(func(context *gin.Context) (any, error) {
		return pairService.List("")
	}))
	api.GET("/list/:prefix", ResponseApiF(func(context *gin.Context) (any, error) {
		var prefix = context.Param("prefix")
		return pairService.List(prefix)
	}))
	api.POST("/save", ResponseApiF(func(context *gin.Context) (any, error) {
		var pair bean.Pair
		err := context.ShouldBindJSON(&pair)
		if err != nil {
			return nil, err
		}
		return pairService.Save(pair)
	}))
	api.POST("/remove", ResponseApiF(func(context *gin.Context) (any, error) {
		var pair bean.Pair
		err := context.ShouldBindJSON(&pair)
		if err != nil {
			return nil, err
		}
		return pairService.RemoveByKey(pair.Key)
	}))
}
