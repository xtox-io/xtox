package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xtox-io/services/handler"
	"github.com/xtox-io/services/server"
)

var (
	configFile string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start services",
	Long:  `Start services of https://xtox.io`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiServer := server.New("web", ":6000", nil, handler.NewAPI())

		errCh := make(chan error, 1)
		apiServer.Start(errCh)

		err := <-errCh

		return err
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&configFile, "config", "c", "/services/config.toml",
		"The config file to use for the services")
}
