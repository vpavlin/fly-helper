package flyhelper

import (
	"log"

	"github.com/spf13/cobra"
)

var version = "v0.0.0"

var rootCmd = &cobra.Command{
	Use:     "flyhelper",
	Version: version,
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.json", "Path to a config file")
	rootCmd.PersistentFlags().String("config-env", "", "Name of the environment variable which contains base64 encoded config file")

}
