package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteVaultCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete removes a context profile",
	Long:    `delete a context with the delete command.`,
	Example: `target vault delete -n example`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		nameToDelete := args[0]

		// Check if the Vault profile with the given name exists and delete it
		if _, exists := c.Vault[args[0]]; exists {
			delete(c.Vault, args[0])
			viper.Set("vault", c.Vault)
			viper.WriteConfig()
			fmt.Printf("Deleted Vault profile '%s'\n", nameToDelete)
		} else {
			fmt.Printf("Vault profile '%s' not found\n", nameToDelete)
		}
	},
}

func init() {
	vaultCmd.AddCommand(deleteVaultCmd)
	//deleteVaultCmd.Flags().StringP("name", "n", "", "Name of the context profile to be deleted")
	//deleteVaultCmd.MarkFlagRequired("name")
}
