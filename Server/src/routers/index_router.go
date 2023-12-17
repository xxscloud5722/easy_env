package routers

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/xxscloud5722/easy_env/server/src/app"
	"log"
	"os"
	"path"
	"time"
)

func (server *Gin) LoadIndex(enable bool) {
	server.GET("/ping", responseApiF(func(context *gin.Context) (any, error) {
		return "pong", nil
	}))
	server.GET("/version", responseApiF(func(context *gin.Context) (any, error) {
		return app.Info().Version, nil
	}))
	server.GET("/datetime", responseApiF(func(context *gin.Context) (any, error) {
		return time.Now().Format(time.DateTime), nil
	}))
	server.GET("/ip", responseApiF(func(context *gin.Context) (any, error) {
		return context.ClientIP(), nil
	}))
	// Admin UI
	if enable {
		server.GET("/admin/*filepath", func(context *gin.Context) {
			filePath := context.Param("filepath")
			filePath = path.Join("./AdminUI", filePath)
			if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
				log.Printf("[filePath]: " + filePath)
				context.File(filePath)
				return
			} else {
				log.Printf("[filePath]: not file " + filePath + "? return index.html")
				context.File(path.Join("./AdminUI", "index.html"))
			}
		})
	}
}
