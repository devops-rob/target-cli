package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	serfendpoint string
	serftoken    string
)
var serfCmd = &cobra.Command{
	Use:   "serf",
	Short: "Manage Serf context profiles ",
	Long: `Manage Serf context profiles.
			Usage: target serf [command]`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"rename",
		"select",
		"update",
	},
	DisableFlagsInUseLine: true,
	Args:                  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serf called")
	},
}

func init() {
	serfCmd.AddCommand(serfCreateCmd)
	serfCmd.AddCommand(deleteCmd)
	serfCmd.AddCommand(renameCmd)
	serfCmd.AddCommand(selectCmd)
	serfCmd.AddCommand(serfUpdateCmd)
	serfCmd.AddCommand(listCmd)
}
