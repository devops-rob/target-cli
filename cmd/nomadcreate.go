package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nomadCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target vault create --name="example" --endpoint="https://example-vault.com:8200"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	// vaultCmd.AddCommand(createCmd)

	nomadCreateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	nomadCreateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-nomad.com:4646")
	nomadCreateCmd.PersistentFlags().StringP("region", "r", "", "set the region of the Nomad server to forward commands to")
	nomadCreateCmd.PersistentFlags().StringP("token", "t", "", "set the  SecretID of an ACL token to use to authenticate API requests with")
	nomadCreateCmd.PersistentFlags().StringP("capath", "p", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	nomadCreateCmd.PersistentFlags().StringP("cacert", "C", "", "set path to a PEM-encoded CA certificate file on the local disk")
	nomadCreateCmd.PersistentFlags().StringP("cert", "c", "", "set path to a PEM-encoded client certificate on the local disk")
	nomadCreateCmd.PersistentFlags().StringP("key", "k", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	nomadCreateCmd.PersistentFlags().StringP("namespace", "N", "", "set Nomad namespace to use for command")

	nomadCreateCmd.MarkPersistentFlagRequired(
		"name",
	)

	nomadCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
