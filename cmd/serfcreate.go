package cmd

import (
	"errors"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serfCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target serf create example --endpoint="https://example-serf.com:7373"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Serf[args[0]] != nil {
			log.Fatal("profile already exists")
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

	serfCreateCmd.PersistentFlags().StringVar(&serfendpoint, "endpoint", "", "set target endpoint details. e.g https://example-serf.com:7373")
	serfCreateCmd.PersistentFlags().StringVar(&serftoken, "token", "", "set serf auth token for this context")

	serfCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
