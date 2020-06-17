package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:     "rename",
	Short:   "rename a vault conext profile",
	Long:    `rename a vault context profile using the rename command`,
	Example: `target vault rename --name="example" --rename="example2"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
	},
}

func init() {
	// vaultCmd.AddCommand(renameCmd)
	renameCmd.PersistentFlags().StringP("name", "n", "", "name of context to be renamed")
	renameCmd.PersistentFlags().StringP("rename", "r", "", "set new name for this context")

	renameCmd.MarkPersistentFlagRequired(
		"name",
	)

	renameCmd.MarkPersistentFlagRequired(
		"rename",
	)

}
