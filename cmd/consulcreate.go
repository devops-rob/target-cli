package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var consulCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target consul create --name="example" --endpoint="https://example-consul.com:8500" --token="s.giqoewbnmdjalkjk"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {

	consulCreateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	consulCreateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-vault.com:8200")
	consulCreateCmd.PersistentFlags().StringP("token", "t", "", "set consul acl token for this context")
	consulCreateCmd.PersistentFlags().StringP("tokenfile", "T", "", "set path to a file containing the API access token for consul")
	consulCreateCmd.PersistentFlags().StringP("capath", "p", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	consulCreateCmd.PersistentFlags().StringP("cacert", "C", "", "set path to a PEM-encoded CA certificate file on the local disk")
	consulCreateCmd.PersistentFlags().StringP("cert", "c", "", "set path to a PEM-encoded client certificate on the local disk")
	consulCreateCmd.PersistentFlags().StringP("key", "k", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")

	consulCreateCmd.MarkPersistentFlagRequired(
		"name",
	)

	consulCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)

}
