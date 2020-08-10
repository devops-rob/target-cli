package cmd

import (
	"errors"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var vaultUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target vault update example --endpoint="https://example2-vault.com:8200" --token="t.loejwikdjuidfhjdi"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Vault[args[0]] == nil {
			log.Fatal("profile does not exists.  Try using the create command")

		}

		v := &Vault{
			Endpoint:  vaultendpoint,
			Token:     vaulttoken,
			CaPath:    vaultcapath,
			CaCert:    vaultcacert,
			Cert:      vaultcert,
			Key:       vaultkey,
			Format:    vaultformat,
			Namespace: vaultnamespace,
		}

		c.Vault[args[0]] = v
		viper.Set("vault", c.Vault)
		viper.WriteConfig()

	},
}

func init() {
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultendpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaulttoken, "token", "", "set vault auth token for this context")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultformat, "format", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultnamespace, "namespace", "", "set vault namespace to use for command")

}
