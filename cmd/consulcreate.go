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
			ConsulEndpoint:  consulendpoint,
			ConsulToken:     consultoken,
			ConsulCaPath:    consulcapath,
			ConsulCaCert:    consulcacert,
			ConsulCert:      consulcert,
			ConsulKey:       consulkey,
			ConsulTokenFile: consultokenfile,
			ConsulNamespace: consulnamespace,
		}

		c.Consul[args[0]] = C

		viper.Set("consul", c.Consul)
		viper.WriteConfig()

	},
}

func init() {

	consulCreateCmd.PersistentFlags().StringVar(&consulendpoint, "endpoint", "", "set target endpoint details. e.g https://example-consul.com:8500")
	consulCreateCmd.PersistentFlags().StringVar(&consultoken, "token", "", "set consul acl token for this context")
	consulCreateCmd.PersistentFlags().StringVar(&consultokenfile, "tokenfile", "", "set path to a file containing the API access token for consul")
	consulCreateCmd.PersistentFlags().StringVar(&consulcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	consulCreateCmd.PersistentFlags().StringVar(&consulnamespace, "namespace", "", "set consul namespace for this context")

	consulCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)

}
