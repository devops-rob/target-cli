package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nomadUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target vault update --name="example" --endpoint="https://example2-nomad.com:8200"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	nomadUpdateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	nomadUpdateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-nomad.com:4646")
	nomadUpdateCmd.PersistentFlags().StringP("region", "r", "", "set the region of the Nomad server to forward commands to")
	nomadUpdateCmd.PersistentFlags().StringP("token", "t", "", "set the  SecretID of an ACL token to use to authenticate API requests with")
	nomadUpdateCmd.PersistentFlags().StringP("capath", "p", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	nomadUpdateCmd.PersistentFlags().StringP("cacert", "C", "", "set path to a PEM-encoded CA certificate file on the local disk")
	nomadUpdateCmd.PersistentFlags().StringP("cert", "c", "", "set path to a PEM-encoded client certificate on the local disk")
	nomadUpdateCmd.PersistentFlags().StringP("key", "k", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	nomadUpdateCmd.PersistentFlags().StringP("namespace", "N", "", "set Nomad namespace to use for command")

	nomadUpdateCmd.MarkPersistentFlagRequired(
		"name",
	)

}
