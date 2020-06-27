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
		fmt.Printf("updating %v target profile with the following values: \n==========================================================\n", serfname)

		if len(serfendpoint) != 0 {
			fmt.Printf("endpoint flag new value: %v\n", serfendpoint)
		}
		if len(serftoken) != 0 {
			fmt.Printf("token flag new value: %v\n", serftoken)
		}
	},
}

func init() {
	serfUpdateCmd.PersistentFlags().StringVar(&serfname, "name", "", "set a profile name for this context")
	serfUpdateCmd.PersistentFlags().StringVar(&serfendpoint, "endpoint", "", "set target endpoint details. e.g https://example-serf.com:7373")
	serfUpdateCmd.PersistentFlags().StringVar(&serftoken, "token", "", "set serf auth token for this context")

	serfUpdateCmd.MarkPersistentFlagRequired(
		"name",
	)
}
