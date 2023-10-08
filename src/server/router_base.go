package server

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//go:embed assets/index.html
var indexFile string

func (server *Server) LoadPing(enable bool) {
	server.GET("/ping", ResponseApiF(func(context *gin.Context) (any, error) {
		return "pong", nil
	}))
	server.GET("/version", ResponseApiF(func(context *gin.Context) (any, error) {
		return "1.2.0", nil
	}))
	server.GET("/datetime", ResponseApiF(func(context *gin.Context) (any, error) {
		return time.Now().Format(time.DateTime), nil
	}))
	// 控制台页面
	if enable {
		server.GET("/admin", func(c *gin.Context) {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(http.StatusOK, indexFile)
		})
	}
}
