package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var nomadCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target nomad create example --endpoint="https://example-nomad.com:8200"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// The following code is to test out CLI behaviour
		fmt.Printf("creating %v target profile with the following values: \n==========================================================\n", args[0])

		fmt.Printf("endpoint flag value: %v\n", nomadendpoint)

		if len(nomadregion) != 0 {
			fmt.Printf("region flag value: %v\n", nomadregion)
		}
		if len(nomadtoken) != 0 {
			fmt.Printf("token flag value: %v\n", nomadtoken)
		}
		if len(nomadcapath) != 0 {
			fmt.Printf("capath flag value: %v\n", nomadcapath)
		}
		if len(nomadcacert) != 0 {
			fmt.Printf("cacert flag value: %v\n", nomadcacert)
		}
		if len(nomadcert) != 0 {
			fmt.Printf("cert flag value: %v\n", nomadcert)
		}
		if len(nomadkey) != 0 {
			fmt.Printf("key flag value: %v\n", nomadkey)
		}
		if len(nomadnamespace) != 0 {
			fmt.Printf("namespace flag value: %v\n", nomadnamespace)
		}
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
