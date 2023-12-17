package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xxscloud5722/easy_env/server/src/bean"
	"github.com/xxscloud5722/easy_env/server/src/service"
)

func (server *Gin) LoadPair() {
	api := server.Group("/pair")
	api.Use(identityFilter)
	var pairService = service.Pair{}
	api.GET("/:key", responseApiF(func(context *gin.Context) (any, error) {
		var key = context.Param("key")
		return pairService.KeyValue(key)
	}))
	api.GET("/list", responseApiF(func(context *gin.Context) (any, error) {
		return pairService.List("")
	}))
	api.GET("/list/:prefix", responseApiF(func(context *gin.Context) (any, error) {
		var prefix = context.Param("prefix")
		return pairService.List(prefix)
	}))
	api.POST("/save", responseApiF(func(context *gin.Context) (any, error) {
		var pair bean.Pair
		err := context.ShouldBindJSON(&pair)
		if err != nil {
			return nil, err
		}
		return pairService.Save(pair)
	}))
	api.POST("/remove", responseApiF(func(context *gin.Context) (any, error) {
		var pair bean.Pair
		err := context.ShouldBindJSON(&pair)
		if err != nil {
			return nil, err
		}
		return pairService.RemoveByKey(pair.Key)
	}))
}
