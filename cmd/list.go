package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listVaultCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault using the list command`,
	Example: `target vault list`,
	Run: func(cmd *cobra.Command, args []string) {

		//var listData [][]string
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Profile Name",
			"Endpoint",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})

		for i, e := range c.Vault {
			data := []string{
				i,
				e.Endpoint,
			}
			table.Append(data)
		}
		table.Render()
	},
}

var listConsulCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Consul using the list command`,
	Example: `target consul list`,
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetHeader([]string{
			"Profile Name",
			"Endpoint",
		})

		for i, e := range c.Consul {
			data := []string{
				i,
				e.ConsulEndpoint,
			}
			table.Append(data)
		}
		table.Render()
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

var listTerraformCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault using the list command`,
	Example: `target terraform list`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("########################################\n##              Terraform             ##\n########################################")
		for i := range c.Terraform {
			fmt.Printf("%s\n", i)
		}
	},
}
