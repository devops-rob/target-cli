package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// updateCmd represents the update command
var consulUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target consul update example --endpoint="https://example2-consul.com:8500" --token="t.loejwikdjuidfhjdi"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		if c.Consul[args[0]] == nil {
			log.Fatal("profile does not exists.  Try using the create command")

		}

		C := &Consul{
			ConsulEndpoint:  consulEndpoint,
			ConsulToken:     consulToken,
			ConsulCaPath:    consulCaPath,
			ConsulCaCert:    consulCaCert,
			ConsulCert:      consulCert,
			ConsulKey:       consulKey,
			ConsulTokenFile: consulTokenFile,
			ConsulNamespace: consulNamespace,
		}

		c.Consul[args[0]] = C

		viper.Set("consul", c.Consul)
		viper.WriteConfig()
	},
}

func init() {
	consulUpdateCmd.PersistentFlags().StringVar(&consulEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	consulUpdateCmd.PersistentFlags().StringVar(&consulToken, "token", "", "set consul acl token for this context")
	consulUpdateCmd.PersistentFlags().StringVar(&consulTokenFile, "tokenfile", "", "set path to a file containing the API access token for consul")
	consulUpdateCmd.PersistentFlags().StringVar(&consulCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	consulUpdateCmd.PersistentFlags().StringVar(&consulCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	consulUpdateCmd.PersistentFlags().StringVar(&consulCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	consulUpdateCmd.PersistentFlags().StringVar(&consulKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	consulUpdateCmd.PersistentFlags().StringVar(&consulNamespace, "namespace", "", "set consul namespace for this context")
}
