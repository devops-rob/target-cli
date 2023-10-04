package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:           "version",
	Short:         "Show current installed version of target-cli",
	Long:          `Show current installed version of target-cli.`,
	Example:       `target vault update example --endpoint="https://example2-vault.com:8200" --token="t.loejwikdjuidfhjdi"`,
	Args:          cobra.NoArgs,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Current Version:", version)
		cmd.Println("")

		return nil
	},
}
