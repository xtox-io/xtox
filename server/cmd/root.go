package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xtox-server",
	Short: "Server of https://xtox.io",
	Long:  `Server (aka backend) is responsible for handling business logic, persistence and integration`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
