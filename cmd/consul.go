package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vaultCmd represents the vault command
var consulCmd = &cobra.Command{
	Use:   "consul",
	Short: "Manage consul context profiles ",
	Long:  `Manage consul context profiles.`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"rename",
		"select",
		"update",
	},
	Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vault called")
	},
}

func init() {
	consulCmd.AddCommand(consulCreateCmd)
	consulCmd.AddCommand(deleteCmd)
	consulCmd.AddCommand(renameCmd)
	consulCmd.AddCommand(selectCmd)
	consulCmd.AddCommand(consulUpdateCmd)
	consulCmd.AddCommand(listCmd)
}
