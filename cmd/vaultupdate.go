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
		fmt.Println("update called")
	},
}

func init() {
	vaultUpdateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	vaultUpdateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultUpdateCmd.PersistentFlags().StringP("token", "t", "", "set vault auth token for this context")
	vaultUpdateCmd.PersistentFlags().StringP("capath", "p", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultUpdateCmd.PersistentFlags().StringP("cacert", "C", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultUpdateCmd.PersistentFlags().StringP("cert", "c", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultUpdateCmd.PersistentFlags().StringP("key", "k", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultUpdateCmd.PersistentFlags().StringP("format", "f", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultUpdateCmd.PersistentFlags().StringP("namespace", "N", "", "set vault namespace to use for command")

	vaultUpdateCmd.MarkPersistentFlagRequired(
		"name",
	)

}
