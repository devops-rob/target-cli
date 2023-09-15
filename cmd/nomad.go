package cmd

import (
	"github.com/spf13/cobra"
)

var (
	nomadendpoint  string
	nomadregion    string
	nomadtoken     string
	nomadcapath    string
	nomadcacert    string
	nomadcert      string
	nomadkey       string
	nomadnamespace string
)

var nomadCmd = &cobra.Command{
	Use:     "nomad",
	Short:   "Manage Nomad context profiles ",
	Long:    `Manage Nomad context profiles.`,
	Example: `target nomad list`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"rename",
		"select",
		"update",
	},
	Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("nomad called")
	//},
}

func init() {
	nomadCmd.AddCommand(nomadCreateCmd)
	//nomadCmd.AddCommand(deleteCmd)
	nomadCmd.AddCommand(renameCmd)
	//nomadCmd.AddCommand(selectCmd)
	nomadCmd.AddCommand(nomadUpdateCmd)
	nomadCmd.AddCommand(listNomadCmd)

}
