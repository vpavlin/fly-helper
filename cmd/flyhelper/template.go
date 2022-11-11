package flyhelper

import (
	"log"

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

		for _, t := range config.Templates {
			err = t.WriteToFile()
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	templateCmd.AddCommand(executeCmd)

	rootCmd.AddCommand(templateCmd)
}
