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
		// The following code is to test out CLI behaviour
		fmt.Printf("creating %v target profile with the following values: \n==========================================================\n", serfname)

		fmt.Printf("endpoint flag value: %v\n", serfendpoint)

		if len(serftoken) != 0 {
			fmt.Printf("token flag value: %v\n", serftoken)
		}

	},
}

func init() {

	serfCreateCmd.PersistentFlags().StringVar(&serfname, "name", "", "set a profile name for this context")
	serfCreateCmd.PersistentFlags().StringVar(&serfendpoint, "endpoint", "", "set target endpoint details. e.g https://example-serf.com:7373")
	serfCreateCmd.PersistentFlags().StringVar(&serftoken, "token", "", "set serf auth token for this context")

	serfCreateCmd.MarkPersistentFlagRequired(
		"name",
	)

	serfCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
