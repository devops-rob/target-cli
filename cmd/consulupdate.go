package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
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
		fmt.Printf("updating %v target profile with the following values: \n==========================================================\n", args[0])

		if len(consulendpoint) != 0 {
			fmt.Printf("endpoint flag new value: %v\n", consulendpoint)
		}
		if len(consultoken) != 0 {
			fmt.Printf("token flag new value: %v\n", consultoken)
		}
		if len(consultokenfile) != 0 {
			fmt.Printf("tokenfile flag new value: %v\n", consultokenfile)
		}
		if len(consulcapath) != 0 {
			fmt.Printf("capath flag new value: %v\n", consulcapath)
		}
		if len(consulcacert) != 0 {
			fmt.Printf("cacert flag new value: %v\n", consulcacert)
		}
		if len(consulcert) != 0 {
			fmt.Printf("cert flag new value: %v\n", consulcert)
		}
		if len(consulkey) != 0 {
			fmt.Printf("key flag new value: %v\n", consulkey)
		}
	},
}

func init() {
	consulUpdateCmd.PersistentFlags().StringVar(&consulendpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	consulUpdateCmd.PersistentFlags().StringVar(&consultoken, "token", "", "set consul acl token for this context")
	consulUpdateCmd.PersistentFlags().StringVar(&consultokenfile, "tokenfile", "", "set path to a file containing the API access token for consul")
	consulUpdateCmd.PersistentFlags().StringVar(&consulcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	consulUpdateCmd.PersistentFlags().StringVar(&consulcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	consulUpdateCmd.PersistentFlags().StringVar(&consulcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	consulUpdateCmd.PersistentFlags().StringVar(&consulkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
}
