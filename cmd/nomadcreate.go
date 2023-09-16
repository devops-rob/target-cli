package cmd

import (
	"errors"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var nomadCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target nomad create example --endpoint="https://example-nomad.com:4646"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Nomad[args[0]] != nil {
			log.Fatal("profile already exists")

		}
		n := &Nomad{
			NomadEndpoint:  nomadendpoint,
			NomadToken:     nomadtoken,
			NomadCaPath:    nomadcapath,
			NomadCaCert:    nomadcacert,
			NomadCert:      nomadcert,
			NomadKey:       nomadkey,
			NomadRegion:    nomadregion,
			NomadNamespace: nomadnamespace,
		}

		c.Nomad[args[0]] = n

		viper.Set("nomad", c.Nomad)
		viper.WriteConfig()

	},
}

func init() {

	nomadCreateCmd.PersistentFlags().StringVar(&nomadendpoint, "endpoint", "", "set target endpoint details. e.g https://example-nomad.com:4646")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadregion, "region", "", "set the region of the Nomad server to forward commands to")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadtoken, "token", "", "set the  SecretID of an ACL token to use to authenticate API requests with")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadnamespace, "namespace", "", "set Nomad namespace to use for command")

	nomadCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
