package cmd

import (
	"fmt"
	"github.com/devops-rob/target-cli/pkg/targetdir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

//var tfVars map[string]string

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

var terraformSetDefaultCmd = &cobra.Command{
	Use:                   "set-default",
	Short:                 "set a default context profile for Terraform ",
	Long:                  `set a default context profile for Terraform.`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.Terraform[args[0]] == nil {
			log.Fatalf("Profile %s not found", args[0])
		}

		context := c.Terraform[args[0]]

		var exportCommandStr []string

		for k, v := range context.Vars {
			command := fmt.Sprintf("export TF_VAR_%s=%s", k, v)
			exportCommandStr = append(exportCommandStr, command)

		}

		commandStr := strings.Join(exportCommandStr, "; ")

		defaultScript := `
#!/bin/bash
` + commandStr

		absolutePath := targetdir.TargetHome() + "/defaults/terraform.sh"

		file, err := os.OpenFile(absolutePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
		_, err = file.WriteString(defaultScript)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("default profile set")
	},
}

func init() {
	terraformCmd.AddCommand(terraformCreateCmd)
	terraformCmd.AddCommand(deleteTerraformCmd)
	terraformCmd.AddCommand(listTerraformCmd)
	terraformCmd.AddCommand(terraformUpdateCmd)
	terraformCmd.AddCommand(selectTerraformCmd)
	terraformCmd.AddCommand(terraformSetDefaultCmd)
}
