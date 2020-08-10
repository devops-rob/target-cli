package cmd

import (
	"errors"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serfUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target serf update example --endpoint="https://example2-serf.com:7373"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Serf[args[0]] == nil {
			log.Fatal("profile does not exists. Try the create command")
		}
		s := &Serf{
			SerfEndopoint: serfendpoint,
			SerfToken:     serftoken,
		}

		c.Serf[args[0]] = s

		viper.Set("serf", c.Serf)
		viper.WriteConfig()
	},
}

func init() {
	serfUpdateCmd.PersistentFlags().StringVar(&serfendpoint, "endpoint", "", "set target endpoint details. e.g https://example-serf.com:7373")
	serfUpdateCmd.PersistentFlags().StringVar(&serftoken, "token", "", "set serf auth token for this context")
}
