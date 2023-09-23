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

var (
	nomadEndpoint  string
	nomadRegion    string
	nomadToken     string
	nomadCaPath    string
	nomadCaCert    string
	nomadCert      string
	nomadKey       string
	nomadNamespace string
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
		"select",
		"update",
		"set-default",
	},
	Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
}

var nomadSetDefaultCmd = &cobra.Command{
	Use:                   "set-default",
	Short:                 "set a default context profile for Nomad ",
	Long:                  `set a default context profile for Nomad.`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.Nomad[args[0]] == nil {
			log.Fatalf("Profile %s not found", args[0])
		}

		context := c.Nomad[args[0]]

		var exportCommandStr []string

		var shellCommandEndpoint string
		endpoint := context.NomadEndpoint
		if endpoint != "" {
			shellCommandEndpoint = fmt.Sprintf("export NOMAD_HTTP_ADDR=%s", endpoint)
			exportCommandStr = append(exportCommandStr, shellCommandEndpoint)
		}

		var shellCommandToken string
		token := context.NomadToken
		if token != "" {
			shellCommandToken = fmt.Sprintf("export NOMAD_HTTP_TOKEN=%s", token)
			exportCommandStr = append(exportCommandStr, shellCommandToken)
		}

		var shellCommandRegion string
		region := context.NomadRegion
		if region != "" {
			shellCommandRegion = fmt.Sprintf("export NOMAD_REGION=%s", region)
			exportCommandStr = append(exportCommandStr, shellCommandRegion)

		}

		var shellCommandCaCert string
		caCert := context.NomadCaCert
		if caCert != "" {
			shellCommandCaCert = fmt.Sprintf("export NOMAD_CACERT=%s", caCert)
			exportCommandStr = append(exportCommandStr, shellCommandCaCert)
		}

		var shellCommandCert string
		cert := context.NomadCert
		if cert != "" {
			shellCommandCert = fmt.Sprintf("export NOMAD_CLIENT_CERT=%s", cert)
			exportCommandStr = append(exportCommandStr, shellCommandCert)
		}

		var shellCommandCaPath string
		caPath := context.NomadCaPath
		if caPath != "" {
			shellCommandCaPath = fmt.Sprintf("export NOMAD_CAPATH=%s", caPath)
			exportCommandStr = append(exportCommandStr, shellCommandCaPath)
		}

		var shellCommandKey string
		key := context.NomadKey
		if key != "" {
			shellCommandKey = fmt.Sprintf("export NOMAD_CLIENT_KEY=%s", key)
			exportCommandStr = append(exportCommandStr, shellCommandKey)
		}

		var shellCommandNameSpace string
		namespace := context.NomadNamespace
		if namespace != "" {
			shellCommandNameSpace = fmt.Sprintf("export NOMAD_NAMESPACE=%s", namespace)
			exportCommandStr = append(exportCommandStr, shellCommandNameSpace)
		}

		commandStr := strings.Join(exportCommandStr, "; ")

		defaultScript := `
#!/bin/bash
` + commandStr

		absolutePath := targetdir.TargetHome() + "/defaults/nomad.sh"

		file, err := os.OpenFile(absolutePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(defaultScript)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("default profile set")
	},
}

func init() {
	nomadCmd.AddCommand(nomadCreateCmd)
	nomadCmd.AddCommand(deleteNomadCmd)
	nomadCmd.AddCommand(nomadSetDefaultCmd)
	nomadCmd.AddCommand(selectNomadCmd)
	nomadCmd.AddCommand(nomadUpdateCmd)
	nomadCmd.AddCommand(listNomadCmd)

}
