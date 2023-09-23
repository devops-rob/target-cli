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
	vaultEndpoint         string
	vaultToken            string
	vaultCaPath           string
	vaultCaCert           string
	vaultCert             string
	vaultKey              string
	vaultFormat           string
	vaultNamespace        string
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

		var shellCommandSkipVerify string
		skipVerify := context.SkipVerify
		if skipVerify != "" {
			shellCommandSkipVerify = fmt.Sprintf("export VAULT_SKIP_VERIFY=%s", skipVerify)
			exportCommandStr = append(exportCommandStr, shellCommandSkipVerify)
		}

		var shellClientTimeout string
		timeout := context.ClientTimeout
		if timeout != "" {
			shellClientTimeout = fmt.Sprintf("export VAULT_CLIENT_TIMEOUT=%s", timeout)
			exportCommandStr = append(exportCommandStr, shellClientTimeout)
		}

		var shellClusterAddr string
		clusterAddr := context.ClusterAddr
		if clusterAddr != "" {
			shellClusterAddr = fmt.Sprintf("export VAULT_CLUSTER_ADDR=%s", clusterAddr)
			exportCommandStr = append(exportCommandStr, shellClusterAddr)
		}

		var shellCommandLicense string
		license := context.License
		if license != "" {
			shellCommandLicense = fmt.Sprintf("export VAULT_LICENSE=%s", license)
			exportCommandStr = append(exportCommandStr, shellCommandLicense)
		}

		var shellCommandLicensePath string
		licensePath := context.LicensePath
		if licensePath != "" {
			shellCommandLicensePath = fmt.Sprintf("export VAULT_LICENSE_PATH=%s", licensePath)
			exportCommandStr = append(exportCommandStr, shellCommandLicensePath)
		}

		var shellCommandLogLevel string
		logLevel := context.LogLevel
		if logLevel != "" {
			shellCommandLogLevel = fmt.Sprintf("export VAULT_LOG_LEVEL=%s", logLevel)
			exportCommandStr = append(exportCommandStr, shellCommandLogLevel)
		}

		var shellCommandMaxRetries string
		maxRetries := context.MaxRetries
		if maxRetries != "" {
			shellCommandMaxRetries = fmt.Sprintf("export VAULT_MAX_RETRIES=%s", maxRetries)
			exportCommandStr = append(exportCommandStr, shellCommandMaxRetries)
		}

		var shellCommandRedirectAddr string
		redirectAddr := context.RedirectAddr
		if redirectAddr != "" {
			shellCommandRedirectAddr = fmt.Sprintf("export VAULT_REDIRECT_ADDR=%s", redirectAddr)
			exportCommandStr = append(exportCommandStr, shellCommandRedirectAddr)
		}

		var shellCommandServerName string
		serverName := context.TlsServerName
		if serverName != "" {
			shellCommandServerName = fmt.Sprintf("export VAULT_TLS_SERVER_NAME=%s", serverName)
			exportCommandStr = append(exportCommandStr, shellCommandServerName)
		}

		var shellCommandCliNoColour string
		cliNoColour := context.CliNoColour
		if cliNoColour != "" {
			shellCommandCliNoColour = fmt.Sprintf("export VAULT_CLI_NO_COLOR=%s", cliNoColour)
			exportCommandStr = append(exportCommandStr, shellCommandCliNoColour)
		}

		var shellCommandRateLimit string
		rateLimit := context.RateLimit
		if rateLimit != "" {
			shellCommandRateLimit = fmt.Sprintf("export VAULT_RATE_LIMIT=%s", rateLimit)
			exportCommandStr = append(exportCommandStr, shellCommandRateLimit)
		}

		var shellCommandSvrLookup string
		svrLookup := context.SvrLookup
		if svrLookup != "" {
			shellCommandSvrLookup = fmt.Sprintf("export VAULT_SRV_LOOKUP=%s", svrLookup)
			exportCommandStr = append(exportCommandStr, shellCommandSvrLookup)
		}

		var shellCommandMfa string
		mfa := context.Mfa
		if svrLookup != "" {
			shellCommandMfa = fmt.Sprintf("export VAULT_MFA=%s", mfa)
			exportCommandStr = append(exportCommandStr, shellCommandMfa)
		}

		var shellCommandHttpProxy string
		httpProxy := context.HttpProxy
		if httpProxy != "" {
			shellCommandHttpProxy = fmt.Sprintf("export VAULT_HTTP_PROXY=%s", httpProxy)
			exportCommandStr = append(exportCommandStr, shellCommandHttpProxy)
		}

		var shellCommandDisableRedirects string
		disableRedirects := context.DisableRedirects
		if disableRedirects != "" {
			shellCommandDisableRedirects = fmt.Sprintf("export VAULT_DISABLE_REDIRECTS=%s", disableRedirects)
			exportCommandStr = append(exportCommandStr, shellCommandDisableRedirects)
		}

		commandStr := strings.Join(exportCommandStr, "; ")

		defaultScript := `
#!/bin/bash
` + commandStr

		absolutePath := targetdir.TargetHome() + "/defaults/vault.sh"

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
	vaultCmd.AddCommand(vaultCreateCmd)
	vaultCmd.AddCommand(deleteVaultCmd)
	vaultCmd.AddCommand(vaultSetDefaultCmd)
	vaultCmd.AddCommand(selectVaultCmd)
	vaultCmd.AddCommand(vaultUpdateCmd)
	vaultCmd.AddCommand(listVaultCmd)

}
