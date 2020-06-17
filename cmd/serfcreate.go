package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serfCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target serf create --name="example" --endpoint="https://example-serf.com:7373"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {

	serfCreateCmd.PersistentFlags().StringP("name", "n", "", "set a profile name for this context")
	serfCreateCmd.PersistentFlags().StringP("endpoint", "e", "", "set target endpoint details. e.g https://example-serf.com:7373")
	serfCreateCmd.PersistentFlags().StringP("token", "t", "", "set serf auth token for this context")

	serfCreateCmd.MarkPersistentFlagRequired(
		"name",
	)

	serfCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
