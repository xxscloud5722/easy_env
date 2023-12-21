package main

import (
	"crypto/rand"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xxscloud5722/easy_env/server/src/server"
	"math/big"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	var serverCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start Easy ENV Server",
		Example: "server -t fa4c7d95a39787f5b62b824b901950e4 -a enable",
		Run: func(cmd *cobra.Command, _args []string) {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			token, err := cmd.Flags().GetString("token")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			if token == "" {
				token = func(length int) string {
					randomString := make([]byte, length)
					for i := range randomString {
						randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
						if err != nil {
							fmt.Println("Error generating random index:", err)
							return ""
						}
						randomString[i] = charset[randomIndex.Int64()]
					}

					return string(randomString)
				}(32)
			}
			admin, err := cmd.Flags().GetBool("admin")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			var args = &server.Args{
				Port:  port,
				Token: token,
				Admin: admin,
			}
			err = server.StartServer(args)
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
		},
	}
	serverCmd.Flags().IntP("port", "p", 8080, "Enable Admin Console")
	serverCmd.Flags().StringP("token", "t", "", "Api Auth AccessToken")
	serverCmd.Flags().BoolP("admin", "a", false, "Enable Admin Console")
	serverCmd.CompletionOptions.HiddenDefaultCmd = true
	err := serverCmd.Execute()
	if err != nil {
		return
	}
}
