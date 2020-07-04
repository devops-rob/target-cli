package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:     "rename [name]",
	Short:   "rename a conext profile",
	Long:    `rename a context profile using the rename command`,
	Example: `target vault rename example --rename="example2"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
	},
}

func init() {
	renameCmd.PersistentFlags().StringP("rename", "r", "", "set new name for this context")

	renameCmd.MarkPersistentFlagRequired(
		"rename",
	)

}
