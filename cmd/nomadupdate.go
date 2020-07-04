package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
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
		// The following code is to test out CLI behaviour
		fmt.Printf("updating %v target profile with the following values: \n==========================================================\n", args[0])

		if len(nomadendpoint) != 0 {
			fmt.Printf("endpoint flag new value: %v\n", nomadendpoint)
		}
		if len(nomadregion) != 0 {
			fmt.Printf("region flag new value: %v\n", nomadregion)
		}
		if len(nomadtoken) != 0 {
			fmt.Printf("token flag new value: %v\n", nomadtoken)
		}
		if len(nomadcapath) != 0 {
			fmt.Printf("capath flag new value: %v\n", nomadcapath)
		}
		if len(nomadcacert) != 0 {
			fmt.Printf("cacert flag new value: %v\n", nomadcacert)
		}
		if len(nomadcert) != 0 {
			fmt.Printf("cert flag new value: %v\n", nomadcert)
		}
		if len(nomadkey) != 0 {
			fmt.Printf("key flag new value: %v\n", nomadkey)
		}
		if len(nomadnamespace) != 0 {
			fmt.Printf("namespace flag new value: %v\n", nomadnamespace)
		}
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
