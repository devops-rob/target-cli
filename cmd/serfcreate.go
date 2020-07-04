package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var serfCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target serf create example --endpoint="https://example-serf.com:7373"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// The following code is to test out CLI behaviour
		fmt.Printf("creating %v target profile with the following values: \n==========================================================\n", args[0])

		fmt.Printf("endpoint flag value: %v\n", serfendpoint)

		if len(serftoken) != 0 {
			fmt.Printf("token flag value: %v\n", serftoken)
		}

	},
}

func init() {

	serfCreateCmd.PersistentFlags().StringVar(&serfendpoint, "endpoint", "", "set target endpoint details. e.g https://example-serf.com:7373")
	serfCreateCmd.PersistentFlags().StringVar(&serftoken, "token", "", "set serf auth token for this context")

	serfCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
