package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewServer() Server {
	return Server{gin.Default()}
}

func (server *Server) StartServer(port int) error {
	err := server.Run(fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}
	return nil
}
