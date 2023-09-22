package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listVaultCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault using the list command`,
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
	Long:    `list all context profiles for Consul using the list command`,
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
	Long:    `list all context profiles for Nomad using the list command`,
	Example: `target nomad list`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("########################################\n##                Nomad               ##\n########################################")
		for i, e := range c.Nomad {
			fmt.Printf("%s:\n  Endpoint: %s\n", i, e.NomadEndpoint)
		}
	},
}

var listBoundaryCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Boundary using the list command`,
	Example: `target boundary list`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("########################################\n##                Boundary               ##\n########################################")
		for i, e := range c.Boundary {
			fmt.Printf("%s:\n  Endpoint: %s\n", i, e.Endpoint)
		}
	},
}
