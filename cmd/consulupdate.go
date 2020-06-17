package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var consulUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target consul update --name="example" --endpoint="https://example2-consul.com:8500" --token="t.loejwikdjuidfhjdi"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	consulUpdateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	consulUpdateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-vault.com:8200")
	consulUpdateCmd.PersistentFlags().StringP("token", "t", "", "set consul acl token for this context")
	consulUpdateCmd.PersistentFlags().StringP("tokenfile", "T", "", "set path to a file containing the API access token for consul")
	consulUpdateCmd.PersistentFlags().StringP("capath", "p", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	consulUpdateCmd.PersistentFlags().StringP("cacert", "C", "", "set path to a PEM-encoded CA certificate file on the local disk")
	consulUpdateCmd.PersistentFlags().StringP("cert", "c", "", "set path to a PEM-encoded client certificate on the local disk")
	consulUpdateCmd.PersistentFlags().StringP("key", "k", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")

	consulUpdateCmd.MarkPersistentFlagRequired(
		"name",
	)

}
