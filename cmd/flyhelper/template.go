package flyhelper

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vpavlin/fly-helper/internal/config"
)

var templateCmd = &cobra.Command{
	Use: "template",
}

var executeCmd = &cobra.Command{
	Use: "execute",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfigFromCommand(cmd)
		if err != nil {
			log.Fatalln(err)
		}

		valuesPath, err := cmd.Flags().GetString("values")
		if err != nil {
			log.Fatalln(err)
		}

		if len(valuesPath) > 0 {
			logrus.Infof("Overriding template values with %s", valuesPath)
			config.Templates.Values = valuesPath
		}

		values, err := config.Templates.LoadValuesFile()
		if err != nil {
			log.Fatalln(err)
		}

		for _, t := range config.Templates.Items {
			logrus.Infof("Processing %s, output to %s", t.Template, t.Output)
			err = t.WriteToFile(values)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	templateCmd.AddCommand(executeCmd)
	executeCmd.Flags().String("values", "", "YAML file containing the values for templating")

	rootCmd.AddCommand(templateCmd)
}
