package flyhelper

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vpavlin/fly-helper/internal/config"
	"github.com/vpavlin/fly-helper/internal/fly"
	"github.com/vpavlin/fly-helper/internal/secrets"
)

var secretsCmd = &cobra.Command{
	Use: "secrets",
}

var pullCmd = &cobra.Command{
	Use: "pull",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfigFromCommand(cmd)
		if err != nil {
			log.Fatalln(err)
		}

		err = config.Secrets.WriteSecrets()
		if err != nil {
			log.Fatalln(err)
		}

	},
}

var pushCmd = &cobra.Command{
	Use: "push",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfigFromCommand(cmd)
		if err != nil {
			log.Fatalln(err)
		}

		fly := fly.NewFly()

		err = config.Push(fly)
		if err != nil {
			if secrets.IsUnchangedErr(err) {
				log.Println(err)
			} else {
				log.Fatalln(err)
			}
		}

		err = config.Secrets.Push(fly, config.AppName)
		if err != nil {
			if secrets.IsUnchangedErr(err) {
				log.Println(err)

			} else {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	secretsCmd.AddCommand(pullCmd)
	secretsCmd.AddCommand(pushCmd)

	rootCmd.AddCommand(secretsCmd)
}
