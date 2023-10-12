package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/nuwa/server.v3/server"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Command() []*cobra.Command {
	var serverCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start Key/Value Server",
		Example: "server -t fa4c7d95a39787f5b62b824b901950e4 -a enable",
		Run: func(cmd *cobra.Command, args []string) {
			port, err := cmd.Flags().GetString("port")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			token, err := cmd.Flags().GetString("token")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			enable, err := cmd.Flags().GetString("admin")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			err = cmdServer(port, token, enable == "enable")
			if err != nil {
				return
			}
		},
	}
	serverCmd.Flags().StringP("port", "p", "", "Enable Admin Console")
	serverCmd.Flags().StringP("token", "t", "", "Api Auth AccessToken")
	serverCmd.Flags().StringP("admin", "a", "", "Enable Admin Console")

	return []*cobra.Command{
		serverCmd,
	}
}

func cmdServer(port, accessToken string, enable bool) error {
	if accessToken == "" {
		accessToken = getToken(32)
	}
	if port == "" {
		port = "8080"
	}
	err := os.Setenv("AccessToken", accessToken)
	if err != nil {
		return err
	}
	gin.SetMode(gin.ReleaseMode)
	var s = server.NewServer()
	s.LoadPair()
	s.LoadScript()
	s.LoadFiles()
	s.LoadPing(enable)
	color.Green(fmt.Sprintf("Server started successfully (Port: %s)..", port))
	color.Blue("AccessToken: " + accessToken)
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	err = s.StartServer(portInt)
	if err != nil {
		return err
	}
	return nil
}

func getToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
