package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xtox-io/client/handler"
	"github.com/xtox-io/client/server"
)

var (
	configFile string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start client",
	Long:  `Start client of https://xtox.io`,
	RunE: func(cmd *cobra.Command, args []string) error {
		webServer := server.New("web", ":5000", nil, handler.NewWeb())
		wsServer := server.New("ws", ":5001", nil, handler.NewWS())

		errCh := make(chan error, 1)
		webServer.Start(errCh)
		wsServer.Start(errCh)

		err := <-errCh

		return err
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&configFile, "config", "c", "/client/config.toml",
		"The config file to use for the client")
}
