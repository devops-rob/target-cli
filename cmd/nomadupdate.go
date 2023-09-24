package cmd

import (
	"errors"
	"fmt"

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
		err := viper.WriteConfig()
		if err != nil {
			return
		}

		fmt.Printf("Updated Nomad profile '%s'\n", args[0])
	},
}

func init() {
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-nomad.com:4646")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadRegion, "region", "", "set the region of the Nomad server to forward commands to")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadToken, "token", "", "set the  SecretID of an ACL token to use to authenticate API requests with")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	nomadUpdateCmd.PersistentFlags().StringVar(&nomadNamespace, "namespace", "", "set Nomad namespace to use for command")

}
