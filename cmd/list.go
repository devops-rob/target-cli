package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault, Nomad, Consul or Serf using the list command`,
	Example: `target vault list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	// vaultCmd.AddCommand(listCmd)
}
