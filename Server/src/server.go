package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxscloud5722/easy_env/server/src/routers"
)

// startServer Run Server.
func startServer(args *Args) error {
	var server = routers.Gin{Engine: gin.Default()}
	server.LoadIndex(args.Admin)
	server.LoadPair()
	err := server.Run(fmt.Sprintf("0.0.0.0:%d", args.Port))
	if err != nil {
		return err
	}
	return nil
}
