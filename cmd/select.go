package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var selectCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target vault select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s context profile selected", args[0])
	},
}

func init() {
	// add flags here
}
