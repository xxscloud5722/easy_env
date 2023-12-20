package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/xxscloud5722/easy_env/server/src/app"
	"github.com/xxscloud5722/easy_env/server/src/bean"
	"github.com/xxscloud5722/easy_env/server/src/service"
	"html/template"
	"net/http"
	"strings"
)

func (server *Gin) LoadFiles() {
	api := server.Group("/files")
	var localFileService = service.LocalFile{
		Root: "./files",
	}
	_ = localFileService.Init()
	api.GET("/*filePath", func(context *gin.Context) {
		var path = strings.TrimLeft(context.Param("filePath"), "/")

		// Path Is Exist
		exist := localFileService.IsExist(path)
		if !exist {
			_ = rendering("404").Execute(context.Writer, gin.H{})
			return
		}
		// Is Dir
		isDir, err := localFileService.IsDir(path)
		if err != nil {
			_ = rendering("500").Execute(context.Writer, gin.H{})
			return
		}
		// Scan Dir Index.
		if *isDir {
			var directory []*bean.DirInfo
			var file []*bean.FileInfo
			directory, file, err = localFileService.ListLocalFile(path)
			if err != nil {
				context.String(http.StatusInternalServerError, "500")
				return
			}
			var paths = strings.Split(path, "/")
			_ = rendering("files").Execute(context.Writer, gin.H{
				"title": "Index of /" + path,
				"directory": lo.Map(directory, func(item *bean.DirInfo, index int) *bean.DirInfo {
					item.Path = "/files/" + path + lo.If(path == "", "").Else("/") + item.Name
					return item
				}),
				"files": lo.Map(file, func(item *bean.FileInfo, index int) *bean.FileInfo {
					item.Path = "/files/" + path + lo.If(path == "", "").Else("/") + item.Name
					return item
				}),
				"isBack": path != "",
				"backPath": "/files/" + lo.IfF(len(paths) > 1, func() string {
					return strings.Join(paths[0:len(paths)-1], "/")
				}).Else(""),
			})
			return
		}

		// Is File Return
		context.File(localFileService.GetFilePath(path))
	})
}

func rendering(key string) *template.Template {
	htmlTemplate := app.Template()[key]
	tmpl, _ := template.New("html").Parse(htmlTemplate)
	return tmpl
}
