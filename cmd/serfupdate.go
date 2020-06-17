package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var serfUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target serf update --name="example" --endpoint="https://example2-serf.com:7373"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	serfUpdateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	serfUpdateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-serf.com:7373")
	serfUpdateCmd.PersistentFlags().StringP("token", "t", "", "set serf auth token for this context")

	serfUpdateCmd.MarkPersistentFlagRequired(
		"name",
	)
}
