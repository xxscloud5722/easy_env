package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nuwa/server.v3/bean"
	"github.com/nuwa/server.v3/service"
	"net/http"
	"os"
	"strings"
)

func (server *Server) LoadScript() {
	api := server.Group("/script")
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
	var scriptService = service.Script{}
	api.GET("/:id", ResponseApiF(func(context *gin.Context) (any, error) {
		var id = context.Param("id")
		return scriptService.Get(id)
	}))
	api.GET("/list", ResponseApiF(func(context *gin.Context) (any, error) {
		var name = context.Query("name")
		var path = context.Query("path")
		return scriptService.List(name, path)
	}))
	api.POST("/save", ResponseApiF(func(context *gin.Context) (any, error) {
		var script bean.Script
		err := context.ShouldBindJSON(&script)
		if err != nil {
			return nil, err
		}
		return nil, scriptService.Save(script)
	}))
	api.POST("/remove", ResponseApiF(func(context *gin.Context) (any, error) {
		var script bean.Script
		err := context.ShouldBindJSON(&script)
		if err != nil {
			return nil, err
		}
		return nil, scriptService.Remove(script.Id)
	}))

	server.GET("/sh/*scriptPath", func(context *gin.Context) {
		var path = context.Param("scriptPath")
		script, err := scriptService.GetByPath(path)
		if err != nil {
			context.String(http.StatusOK, "")
		}
		if script == nil {
			context.String(http.StatusOK, "")
		}
		context.String(http.StatusOK, script.Value)
	})
}
