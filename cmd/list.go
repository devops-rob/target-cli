package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listVaultCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault, Nomad, Consul or Serf using the list command`,
	Example: `target vault list`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("########################################\n##                Vault               ##\n########################################")
		for i, e := range c.Vault {
			fmt.Printf("%s:\n  Endpoint: %s\n", i, e.Endpoint)
		}
	},
}

var listConsulCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault, Nomad, Consul or Serf using the list command`,
	Example: `target consul list`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("########################################\n##                Consul              ##\n########################################")
		for i, e := range c.Consul {
			fmt.Printf("%s:\n  Endpoint: %s\n", i, e.ConsulEndpoint)
		}
	},
}

var listNomadCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault, Nomad, Consul or Serf using the list command`,
	Example: `target nomad list`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("########################################\n##                Nomad               ##\n########################################")
		for i, e := range c.Nomad {
			fmt.Printf("%s:\n  Endpoint: %s\n", i, e.NomadEndpoint)
		}
	},
}
