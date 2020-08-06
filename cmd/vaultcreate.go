package cmd

import (
	"errors"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vaultCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target vault create example --endpoint="https://example-vault.com:8200" --token="s.giqoewbnmdjalkjk"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.Vault[args[0]] != nil {
			log.Fatal("profile already exists")

		}
		v := &Vault{
			Endpoint: vaultendpoint,
			Token:    vaulttoken,
		}

		c.Vault[args[0]] = v

		// fmt.Printf("%+v\n", c)
		viper.Set("vault", c.Vault)
		// fmt.Printf("%+v\n", viper.AllSettings())
		viper.WriteConfig()

	},
}

func init() {

	vaultCreateCmd.PersistentFlags().StringVar(&vaultendpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultCreateCmd.PersistentFlags().StringVar(&vaulttoken, "token", "", "set vault auth token for this context")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultformat, "format", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultCreateCmd.PersistentFlags().StringVar(&vaultnamespace, "namespace", "", "set vault namespace to use for command")

	vaultCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)

}
