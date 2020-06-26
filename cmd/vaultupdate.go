package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var vaultUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target vault update --name="example" --endpoint="https://example2-vault.com:8200" --token="t.loejwikdjuidfhjdi"`,
	Run: func(cmd *cobra.Command, args []string) {
		// The following code is to test out CLI behaviour
		fmt.Printf("updating %v target profile with the following values: \n==========================================================\n", vaultname)

		fmt.Printf("endpoint flag new value: %v\n", vaultendpoint)

		if len(vaultformat) != 0 {
			fmt.Printf("format flag new value: %v\n", vaultformat)
		}
		if len(vaulttoken) != 0 {
			fmt.Printf("token flag new value: %v\n", vaulttoken)
		}
		if len(vaultcapath) != 0 {
			fmt.Printf("capath flag new value: %v\n", vaultcapath)
		}
		if len(vaultcacert) != 0 {
			fmt.Printf("cacert flag new value: %v\n", vaultcacert)
		}
		if len(vaultcert) != 0 {
			fmt.Printf("cert flag new value: %v\n", vaultcert)
		}
		if len(vaultkey) != 0 {
			fmt.Printf("key flag new value: %v\n", vaultkey)
		}
		if len(vaultnamespace) != 0 {
			fmt.Printf("namespace flag new value: %v\n", vaultnamespace)
		}
	},
}

func init() {
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultname, "name", "", "set a profile name for this context")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultendpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaulttoken, "token", "", "set vault auth token for this context")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultformat, "format", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultnamespace, "namespace", "", "set vault namespace to use for command")

	vaultUpdateCmd.MarkPersistentFlagRequired(
		"name",
	)

}
