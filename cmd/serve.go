package cmd

import (
	"github.com/chrootlogin/stato/pkg/webserver"
	"github.com/spf13/cobra"
)

func init() {
	var listenAddress string

	serveGenerate := &cobra.Command{
		Use:   "serve",
		Short: "Webserver",
		Long: `serve is for running a webserver`,
		Run: func(cmd *cobra.Command, args []string) {
			// run webserver
			webserver.New(workDir).Run(listenAddress)
		},
	}

	// flag for listen address
	serveGenerate.Flags().StringVarP(&listenAddress, "listen", "l", ":1234", "listen address of the webserver")

	rootCmd.AddCommand(serveGenerate)
}