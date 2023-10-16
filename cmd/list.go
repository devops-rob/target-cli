package cmd

import (
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

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetHeader([]string{
			"Profile Name",
			"Endpoint",
		})

		for i, e := range c.Nomad {
			data := []string{
				i,
				e.NomadEndpoint,
			}
			table.Append(data)
		}
		table.Render()
	},
}

var listBoundaryCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Boundary using the list command`,
	Example: `target boundary list`,
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetHeader([]string{
			"Profile Name",
			"Endpoint",
		})

		for i, e := range c.Boundary {
			data := []string{
				i,
				e.Endpoint,
			}
			table.Append(data)
		}
		table.Render()
	},
}

var listTerraformCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for your chosen tool",
	Long:    `list all context profiles for Vault using the list command`,
	Example: `target terraform list`,
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetHeader([]string{
			"Profile Name",
		})

		for i := range c.Terraform {
			data := []string{
				i,
			}
			table.Append(data)
		}
		table.Render()
	},
}
