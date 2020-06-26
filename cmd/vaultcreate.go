package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var vaultCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target vault create --name="example" --endpoint="https://example-vault.com:8200" --token="s.giqoewbnmdjalkjk"`,
	Run: func(cmd *cobra.Command, args []string) {
		// The following code is to test out CLI behaviour
		fmt.Printf("creating %v target profile with the following values: \n==========================================================\n", vaultname)

		fmt.Printf("endpoint flag value: %v\n", vaultendpoint)

		if len(vaultformat) != 0 {
			fmt.Printf("format flag value: %v\n", vaultformat)
		}
		if len(vaulttoken) != 0 {
			fmt.Printf("token flag value: %v\n", vaulttoken)
		}
		if len(vaultcapath) != 0 {
			fmt.Printf("capath flag value: %v\n", vaultcapath)
		}
		if len(vaultcacert) != 0 {
			fmt.Printf("cacert flag value: %v\n", vaultcacert)
		}
		if len(vaultcert) != 0 {
			fmt.Printf("cert flag value: %v\n", vaultcert)
		}
		if len(vaultkey) != 0 {
			fmt.Printf("key flag value: %v\n", vaultkey)
		}
		if len(vaultnamespace) != 0 {
			fmt.Printf("namespace flag value: %v\n", vaultnamespace)
		}
	},
}

func init() {

	vaultCreateCmd.PersistentFlags().StringVar(&vaultname, "name", "", "set a profile name for this context")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultendpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultCreateCmd.PersistentFlags().StringVar(&vaulttoken, "token", "", "set vault auth token for this context")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultcapath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultcacert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultcert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultkey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultformat, "format", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultCreateCmd.PersistentFlags().StringVar(&vaultnamespace, "namespace", "", "set vault namespace to use for command")

	vaultCreateCmd.MarkPersistentFlagRequired(
		"name",
	)

	vaultCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
