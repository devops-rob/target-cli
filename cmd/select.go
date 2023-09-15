package cmd

import (
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"strings"
)

// switchCmd represents the switch command
var selectVaultCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target vault select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Vault[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Vault[args[0]]

		var exportCommandStr []string

		var shellCommandEndpoint string
		endpoint := context.Endpoint
		if endpoint != "" {
			shellCommandEndpoint = fmt.Sprintf("export VAULT_ADDR=%s", endpoint)
			exportCommandStr = append(exportCommandStr, shellCommandEndpoint)
		}

		var shellCommandToken string
		token := context.Token
		if token != "" {
			shellCommandToken = fmt.Sprintf("export VAULT_TOKEN=%s", token)
			exportCommandStr = append(exportCommandStr, shellCommandToken)
		}

		var shellCommandNamespace string
		namespace := context.Namespace
		if namespace != "" {
			shellCommandNamespace = fmt.Sprintf("export VAULT_NAMESPACE=%s", namespace)
			exportCommandStr = append(exportCommandStr, shellCommandNamespace)

		}

		var shellCommandCaCert string
		caCert := context.CaCert
		if caCert != "" {
			shellCommandCaCert = fmt.Sprintf("export VAULT_CACERT=%s", caCert)
			exportCommandStr = append(exportCommandStr, shellCommandCaCert)
		}

		var shellCommandCert string
		cert := context.Cert
		if cert != "" {
			shellCommandCert = fmt.Sprintf("export VAULT_CLIENT_CERT=%s", cert)
			exportCommandStr = append(exportCommandStr, shellCommandCert)
		}

		var shellCommandCaPath string
		caPath := context.CaPath
		if caPath != "" {
			shellCommandCaPath = fmt.Sprintf("export VAULT_CAPATH=%s", caPath)
			exportCommandStr = append(exportCommandStr, shellCommandCaPath)
		}

		var shellCommandKey string
		key := context.Key
		if key != "" {
			shellCommandKey = fmt.Sprintf("export VAULT_CLIENT_KEY=%s", key)
			exportCommandStr = append(exportCommandStr, shellCommandKey)
		}

		var shellCommandFormat string
		format := context.Format
		if format != "" {
			shellCommandFormat = fmt.Sprintf("export VAULT_FORMAT=%s", format)
			exportCommandStr = append(exportCommandStr, shellCommandFormat)
		}

		commandStr := strings.Join(exportCommandStr, "; ")
		fmt.Println(commandStr)
	},
}

func init() {
	vaultCmd.AddCommand(selectVaultCmd)
}
