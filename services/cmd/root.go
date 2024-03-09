package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xtox-services",
	Short: "Services of https://xtox.io",
	Long:  `Services (aka lambda) is responsible for user functions`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
