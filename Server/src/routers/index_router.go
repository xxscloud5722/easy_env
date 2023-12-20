package routers

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/xxscloud5722/easy_env/server/src/app"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
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
			localFile, err := app.AdminDir.Open("admin" + filePath)
			if err == nil {
				var fileInfo fs.FileInfo
				fileInfo, err = localFile.Stat()
				if err != nil {
					return
				}
				if !fileInfo.IsDir() {
					var bytes []byte
					bytes, err = app.AdminDir.ReadFile("admin" + filePath)
					if err != nil {
						return
					}
					context.Data(http.StatusOK, mime.TypeByExtension(strings.ToLower(filepath.Ext(fileInfo.Name()))), bytes)
					return
				}
			}
			bytes, err := app.AdminDir.ReadFile("admin/index.html")
			if err != nil {
				return
			}
			context.Data(http.StatusOK, "text/html", bytes)
		})
	}
}
