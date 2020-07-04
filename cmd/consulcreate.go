package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var consulCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target consul create example --endpoint="https://example-consul.com:8500" --token="s.giqoewbnmdjalkjk"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("creating %v target profile with the following values: \n==========================================================\n", args[0])

		fmt.Printf("endpoint flag value: %v\n", consulendpoint)

		if len(consultoken) != 0 {
			fmt.Printf("token flag value: %v\n", consultoken)
		}
		if len(consultokenfile) != 0 {
			fmt.Printf("tokenfile flag value: %v\n", consultokenfile)
		}
		if len(consulcapath) != 0 {
			fmt.Printf("capath flag value: %v\n", consulcapath)
		}
		if len(consulcacert) != 0 {
			fmt.Printf("cacert flag value: %v\n", consulcacert)
		}
		if len(consulcert) != 0 {
			fmt.Printf("cert flag value: %v\n", consulcert)
		}
		if len(consulkey) != 0 {
			fmt.Printf("key flag value: %v\n", consulkey)
		}
	},
}

func init() {

	consulCreateCmd.PersistentFlags().StringVar(&consulendpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	consulCreateCmd.PersistentFlags().StringVar(&consultoken, "token", "", "set consul acl token for this context")
	consulCreateCmd.PersistentFlags().StringVar(&consultokenfile, "tokenfile", "", "set path to a file containing the API access token for consul")
	consulCreateCmd.PersistentFlags().StringVar(&consulcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	consulCreateCmd.PersistentFlags().StringVar(&consulkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")

	consulCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)

}
