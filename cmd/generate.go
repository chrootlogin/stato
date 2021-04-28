package cmd

import (
	"github.com/chrootlogin/stato/pkg/stato"
	"github.com/spf13/cobra"
)

func init() {
	cmdGenerate := &cobra.Command{
		Use:   "gen",
		Short: "Generate pages",
		Long: `generate is for generating all pages`,
		Run: func(cmd *cobra.Command, args []string) {
			// build all pages
			stato.Load(workDir).BuildAll()
		},
	}

	rootCmd.AddCommand(cmdGenerate)
}