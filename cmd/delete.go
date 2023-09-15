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

var deleteConsulCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete removes a context profile",
	Long:    `delete a context with the delete command.`,
	Example: `target consul delete -n example`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		nameToDelete := args[0]

		if _, exists := c.Consul[args[0]]; exists {
			delete(c.Consul, args[0])
			viper.Set("consul", c.Consul)
			viper.WriteConfig()
			fmt.Printf("Deleted Consul profile '%s'\n", nameToDelete)
		} else {
			fmt.Printf("Consul profile '%s' not found\n", nameToDelete)
		}
	},
}

var deleteNomadCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete removes a context profile",
	Long:    `delete a context with the delete command.`,
	Example: `target nomad delete -n example`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		nameToDelete := args[0]

		if _, exists := c.Nomad[args[0]]; exists {
			delete(c.Nomad, args[0])
			viper.Set("nomad", c.Nomad)
			viper.WriteConfig()
			fmt.Printf("Deleted Nomad profile '%s'\n", nameToDelete)
		} else {
			fmt.Printf("Nomad profile '%s' not found\n", nameToDelete)
		}
	},
}
