package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Args struct {
	Port  int    // server Port
	Token string // set Token
	Admin bool   // enable Admin
}

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
			admin, err := cmd.Flags().GetBool("admin")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			var args = &Args{
				Port:  port,
				Token: token,
				Admin: admin,
			}
			err = startServer(args)
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
		},
	}
	serverCmd.Flags().IntP("port", "p", 8080, "Enable Admin Console")
	serverCmd.Flags().StringP("token", "t", "", "Api Auth AccessToken")
	serverCmd.Flags().BoolP("admin", "a", false, "Enable Admin Console")

	err := serverCmd.Execute()
	if err != nil {
		return
	}
}
