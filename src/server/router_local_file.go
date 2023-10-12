package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nuwa/server.v3/service"
	"net/http"
	"strings"
)

func (server *Server) LoadFiles() {
	api := server.Group("/files")
	var localFileService = service.LocalFile{
		Root: "./files",
	}
	_ = localFileService.Init()
	api.GET("/*filePath", func(context *gin.Context) {
		var path = strings.TrimLeft(context.Param("filePath"), "/")
		// 路径不明确
		if path == "" {
			context.String(http.StatusNotFound, "404 page not found")
			return
		}

		// 文件是否存在
		exist := localFileService.IsExist(path)
		if !exist {
			context.String(http.StatusNotFound, "404 page not found")
			return
		}
		// 如果是目录
		isDir, err := localFileService.IsDir(path)
		if err != nil {
			context.String(http.StatusInternalServerError, "500 server error")
			return
		}
		if *isDir {
			var directory []string
			var file []string
			directory, file, err = localFileService.ListLocalFile(path)
			if err != nil {
				context.String(http.StatusInternalServerError, "500")
				return
			}
			context.JSON(http.StatusOK, struct {
				directory []string
				file      []string
			}{
				directory: directory,
				file:      file,
			})
			return
		}

		// 如果是文件
		context.File(localFileService.GetFilePath(path))
	})
}
