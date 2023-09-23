package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
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
			NomadEndpoint:  nomadEndpoint,
			NomadToken:     nomadToken,
			NomadCaPath:    nomadCaPath,
			NomadCaCert:    nomadCaCert,
			NomadCert:      nomadCert,
			NomadKey:       nomadKey,
			NomadRegion:    nomadRegion,
			NomadNamespace: nomadNamespace,
		}

		c.Nomad[args[0]] = n

		viper.Set("nomad", c.Nomad)
		viper.WriteConfig()

	},
}

func init() {

	nomadCreateCmd.PersistentFlags().StringVar(&nomadEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-nomad.com:4646")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadRegion, "region", "", "set the region of the Nomad server to forward commands to")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadToken, "token", "", "set the  SecretID of an ACL token to use to authenticate API requests with")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	nomadCreateCmd.PersistentFlags().StringVar(&nomadNamespace, "namespace", "", "set Nomad namespace to use for command")

	nomadCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
