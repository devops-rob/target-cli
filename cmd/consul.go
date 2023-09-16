package cmd

import (
	"github.com/spf13/cobra"
)

var (
	consulendpoint  string
	consultoken     string
	consultokenfile string
	consulcapath    string
	consulcacert    string
	consulcert      string
	consulkey       string
	consulnamespace string
)

var consulCmd = &cobra.Command{
	Use:   "consul",
	Short: "Manage consul context profiles ",
	Long:  `Manage consul context profiles.`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"select",
		"update",
	},
	Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("consul called")
	//},
}

func init() {
	consulCmd.AddCommand(consulCreateCmd)
	//consulCmd.AddCommand(deleteConsulCmd)
	//consulCmd.AddCommand(renameCmd)
	consulCmd.AddCommand(selectConsulCmd)
	consulCmd.AddCommand(consulUpdateCmd)
	consulCmd.AddCommand(listConsulCmd)
}
