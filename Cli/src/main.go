package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xxscloud5722/easy_env/cli/src/environment"
)

func main() {
	var cliCmd = &cobra.Command{
		Use:     "cli",
		Short:   "Easy ENV Client",
		Example: "cli list",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			token, err := cmd.Flags().GetString("token")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			url, err := cmd.Flags().GetString("url")
			if err != nil {
				color.Red(fmt.Sprint(err))
				return
			}
			environment.SetArgs(url, token)
		},
	}
	cliCmd.PersistentFlags().StringP("token", "t", "", "Api Auth AccessToken")
	cliCmd.PersistentFlags().StringP("url", "u", "http://127.0.0.1:8080", "Server URL")
	cliCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all variables",
		Run: func(cmd *cobra.Command, args []string) {
			err := environment.Print(lo.IfF(len(args) > 0, func() string { return args[0] }).Else(""))
			if err != nil {
				color.Red(fmt.Sprint(err))
			}
		},
	})
	cliCmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get a single variable",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 0 {
				return
			}
			err := environment.PrintByKey(args[0])
			if err != nil {
				color.Red(fmt.Sprint(err))
			}
		},
	})
	cliCmd.AddCommand(&cobra.Command{
		Use:   "push",
		Short: "Add variables",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				return
			}
			err := environment.Push(args[0], args[1], lo.IfF(len(args) > 2, func() string { return args[2] }).Else(""))
			if err != nil {
				color.Red(fmt.Sprint(err))
			}
		},
	})
	cliCmd.AddCommand(&cobra.Command{
		Use:   "remove",
		Short: "Remove variables",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 0 {
				return
			}
			err := environment.Remove(args[0])
			if err != nil {
				color.Red(fmt.Sprint(err))
			}
		},
	})
	cliCmd.CompletionOptions.HiddenDefaultCmd = true
	err := cliCmd.Execute()
	if err != nil {
		return
	}
}
