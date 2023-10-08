package main

import (
	"github.com/nuwa/server.v3/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "server"}
	for _, it := range cmd.Command() {
		rootCmd.AddCommand(it)
	}
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
