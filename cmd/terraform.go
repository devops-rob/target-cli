package cmd

import "github.com/spf13/cobra"

var tfVars map[string]string

var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "Manage Terraform context profiles ",
	Long:  `Manage Terraform context profiles.`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"select",
		"update",
		"set-default",
	},
	DisableFlagsInUseLine: true,
}

func init() {
	terraformCmd.AddCommand(terraformCreateCmd)
	terraformCmd.AddCommand(deleteTerraformCmd)
	terraformCmd.AddCommand(listTerraformCmd)
}
