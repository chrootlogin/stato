package cmd

import (
	"github.com/chrootlogin/stato/pkg/utils/consts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	// used for flags.
	cfgFile string
	workDir string
	logLevel string

	// root command
	rootCmd = &cobra.Command{
		Use:   "stato",
		Short: "A static page generator",
		Long: `Stato allows you to render template files to a webpage.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	// flags for config
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "", "log level")
	rootCmd.PersistentFlags().StringVar(&workDir, "workdir", "", "working directory")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// Cli interface for stato
/*func CliInterface() *cobra.Command {
	cobra.OnInitialize(initConfig)

	rootCmd := &cobra.Command{
		Use:   "stato",
		Short: "A static page generator",
		Long: `Stato allows you to render template files to a webpage.`,
	}

	// flags for config
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "", "log level")
	rootCmd.PersistentFlags().StringVar(&workDir, "workdir", "", "working directory")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	// add commands
	rootCmd.AddCommand(cmdGenerate)
	rootCmd.AddCommand(serveGenerate)

	return rootCmd
}*/

func initConfig() {
	if workDir == "" {
		currentWorkingDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		workDir = currentWorkingDirectory
	}

	if cfgFile == "" {
		cfgFile = consts.StatoDefaultCfgFile
	}

	if logLevel == "" {
		log.SetLevel(log.InfoLevel)
	} else {
		switch logLevel {
		case "debug":
			log.SetLevel(log.DebugLevel)
			break
		case "info":
			log.SetLevel(log.InfoLevel)
			break
		case "warning":
			log.SetLevel(log.WarnLevel)
			break
		case "error":
			log.SetLevel(log.ErrorLevel)
			break
		default:
			log.SetLevel(log.InfoLevel)
			break
		}
	}

	// enable configuration
	viper.SetConfigFile(filepath.Join(workDir, cfgFile))
	log.WithField("path", workDir).Info("setting work directory")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.WithField("path",viper.ConfigFileUsed()).Info("loading config file")
	}
}