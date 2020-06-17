package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var selectCmd = &cobra.Command{
	Use:     "select",
	Short:   "select to a context",
	Long:    `select a context to use with the select command.`,
	Example: `target vault switch --name="example"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch called")
	},
}

func init() {
	// vaultCmd.AddCommand(selectCmd)
	selectCmd.PersistentFlags().StringP("name", "n", "", "name of context to switch to")

	selectCmd.MarkPersistentFlagRequired(
		"name",
	)

}
