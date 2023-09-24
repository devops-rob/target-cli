package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var terraformUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target terraform update example --var rob=theauthor`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Terraform[args[0]] == nil {
			log.Fatal("profile does not exists.  Try using the create command")

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

		c.Terraform[args[0]] = t
		viper.Set("terraform", c.Terraform)
		err := viper.WriteConfig()
		if err != nil {
			return
		}

	},
}

func init() {
	terraformUpdateCmd.PersistentFlags().StringSliceVarP(&tfVarFlag, "var", "v", []string{}, "set a terraform variable in a key/value pair, e.g name=rob. Can be specified multiple times")

	err := terraformCreateCmd.MarkPersistentFlagRequired(
		"var",
	)
	if err != nil {
		return
	}

}
