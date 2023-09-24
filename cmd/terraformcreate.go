package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var tfVarFlag []string

var terraformCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target terraform create example --var "project_id=1234567" --var "vault_addr=https://prd-vault.target:8200"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.Terraform[args[0]] != nil {
			log.Fatal("profile already exists")

		}

		result := make(map[string]string)
		for _, item := range tfVarFlag {
			parts := strings.SplitN(item, "=", 2)
			if len(parts) == 2 {
				result[parts[0]] = parts[1]
			}
		}

		t := &Terraform{
			Vars: result,
		}

		c.Terraform[args[0]] = t

		viper.Set("terraform", c.Terraform)
		err := viper.WriteConfig()
		if err != nil {
			return
		}
		fmt.Printf("Created Terraform profile '%s'\n", args[0])

	},
}

func init() {

	//terraformCreateCmd.Flags().StringSlice("var", map[string]string, "set a terraform variable in a key/value pair, e.g name=rob. Can be specified multiple times")
	terraformCreateCmd.PersistentFlags().StringSliceVarP(&tfVarFlag, "var", "v", []string{}, "set a terraform variable in a key/value pair, e.g name=rob. Can be specified multiple times")

	err := terraformCreateCmd.MarkPersistentFlagRequired(
		"var",
	)
	if err != nil {
		return
	}

}
