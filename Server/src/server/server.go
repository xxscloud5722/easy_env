package server

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/xxscloud5722/easy_env/server/src/routers"
	"os"
)

type Args struct {
	Port  int    // server Port
	Token string // set Token
	Admin bool   // enable Admin
}

// StartServer Run Server.
func StartServer(args *Args) error {
	err := os.Setenv("AccessToken", args.Token)
	if err != nil {
		return err
	}
	color.Blue("[EnvServer] Token: " + args.Token)
	gin.SetMode(gin.ReleaseMode)
	var server = routers.Gin{Engine: gin.Default()}
	server.LoadIndex(args.Admin)
	server.LoadPair()
	server.LoadFiles()
	err = server.Run(fmt.Sprintf("0.0.0.0:%d", args.Port))
	if err != nil {
		return err
	}
	return nil
}
