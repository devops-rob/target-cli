package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var consulCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target consul create example --endpoint="https://example-consul.com:8500" --token="a-consul-acl-token"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Consul[args[0]] != nil {
			log.Fatal("profile already exists")

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

	consulCreateCmd.PersistentFlags().StringVar(&consulEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-consul.com:8500")
	consulCreateCmd.PersistentFlags().StringVar(&consulToken, "token", "", "set consul acl token for this context")
	consulCreateCmd.PersistentFlags().StringVar(&consulTokenFile, "tokenfile", "", "set path to a file containing the API access token for consul")
	consulCreateCmd.PersistentFlags().StringVar(&consulCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	consulCreateCmd.PersistentFlags().StringVar(&consulNamespace, "namespace", "", "set consul namespace for this context")

	consulCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)

}
