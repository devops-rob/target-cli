package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete removes a context profile",
	Long:    `delete a  context with the delete command.`,
	Example: `target vault delete --name="example"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	deleteCmd.PersistentFlags().StringP("name", "n", "", "name of context to be deleted")

	deleteCmd.MarkPersistentFlagRequired(
		"name",
	)
}
