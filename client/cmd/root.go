package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xtox-client",
	Short: "Client of https://xtox.io",
	Long:  `Client (aka frontend) is responsible for accepting requests, delegating to server and rendering the response`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
