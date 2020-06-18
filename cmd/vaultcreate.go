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
		fmt.Println("create called")
	},
}

func init() {

	vaultCreateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	vaultCreateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultCreateCmd.PersistentFlags().StringP("token", "t", "", "set vault auth token for this context")
	vaultCreateCmd.PersistentFlags().StringP("capath", "p", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultCreateCmd.PersistentFlags().StringP("cacert", "C", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultCreateCmd.PersistentFlags().StringP("cert", "c", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultCreateCmd.PersistentFlags().StringP("key", "k", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultCreateCmd.PersistentFlags().StringP("format", "f", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultCreateCmd.PersistentFlags().StringP("namespace", "N", "", "set vault namespace to use for command")

	vaultCreateCmd.MarkPersistentFlagRequired(
		"name",
	)

	vaultCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
