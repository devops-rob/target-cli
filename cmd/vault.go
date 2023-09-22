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
	vaultendpoint  string
	vaulttoken     string
	vaultcapath    string
	vaultcacert    string
	vaultcert      string
	vaultkey       string
	vaultformat    string
	vaultnamespace string
	// flags to add
	vaultSkipVerify       string
	vaultClientTimeout    string
	vaultClusterAddr      string
	vaultLicense          string
	vaultLicensePath      string
	vaultLogLevel         string
	vaultMaxRetries       string
	vaultRedirectAddr     string
	vaultTlsServerName    string
	vaultCliNoColour      string
	vaultRateLimit        string
	vaultSvrLookup        string
	vaultMfa              string
	vaultHttpProxy        string
	vaultDisableRedirects string
)

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Manage Vault context profiles ",
	Long:  `Manage Vault context profiles.`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"select",
		"update",
		"set-default",
	},
	//Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
}

var vaultSetDefaultCmd = &cobra.Command{
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

		if c.Vault[args[0]] == nil {
			log.Fatalf("Profile %s not found", args[0])
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

		//var shellCommandRegion string
		//region := context.NomadRegion
		//if region != "" {
		//	shellCommandRegion = fmt.Sprintf("export NOMAD_REGION=%s", region)
		//	exportCommandStr = append(exportCommandStr, shellCommandRegion)
		//
		//}

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

		var shellCommandNameSpace string
		namespace := context.Namespace
		if namespace != "" {
			shellCommandNameSpace = fmt.Sprintf("export VAULT_NAMESPACE=%s", namespace)
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
	vaultCmd.AddCommand(vaultCreateCmd)
	vaultCmd.AddCommand(deleteVaultCmd)
	//vaultCmd.AddCommand(renameCmd)
	vaultCmd.AddCommand(selectVaultCmd)
	vaultCmd.AddCommand(vaultUpdateCmd)
	vaultCmd.AddCommand(listVaultCmd)

}
