package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var nomadUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target nomad update example --endpoint="https://example2-nomad.com:8200"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Nomad[args[0]] == nil {
			log.Fatal("profile does not exists.  Try using the create command")

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
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadendpoint, "endpoint", "", "set target endpoint details. e.g https://example-nomad.com:4646")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadregion, "region", "", "set the region of the Nomad server to forward commands to")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadtoken, "token", "", "set the  SecretID of an ACL token to use to authenticate API requests with")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadnamespace, "namespace", "", "set Nomad namespace to use for command")

}
