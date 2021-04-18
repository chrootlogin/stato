package main

import (
	"github.com/chrootlogin/stato/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

func configureViper() {

}

func runCli() {
	c := dig.New()

	if err := c.Provide(cmd.CliInterface); err != nil {
		log.Fatal("Error loading root command", err)
	}

	c.Invoke(func(cliInterface *cobra.Command) {
		if err := cliInterface.Execute(); err != nil {
			log.Fatal("Error loading cli interface", err)
		}
	})
}

func main() {
	configureViper()
	runCli()
}