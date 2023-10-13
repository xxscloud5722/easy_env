package server

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"time"
)

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
		server.GET("/admin/*filepath", func(context *gin.Context) {
			filePath := context.Param("filepath")
			filePath = path.Join("E:\\code\\xxscloud\\github\\bpp\\server\\AdminUI\\dist", filePath)
			if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
				context.File(filePath)
				return
			} else {
				context.File(path.Join("./AdminUI", "index.html"))
			}
		})
	}
}
