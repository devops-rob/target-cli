package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// selectVaultCmd represents the switch command for Vault
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

		exportCommands := []string{}

		if context.Endpoint != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_ADDR=%s", context.Endpoint))
		}

		if context.Token != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_TOKEN=%s", context.Token))
		}

		if context.Namespace != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_NAMESPACE=%s", context.Namespace))
		}

		if context.CaCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CACERT=%s", context.CaCert))
		}

		if context.Cert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_CERT=%s", context.Cert))
		}

		if context.CaPath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CAPATH=%s", context.CaPath))
		}

		if context.Key != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_KEY=%s", context.Key))
		}

		if context.Format != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_FORMAT=%s", context.Format))
		}

		if context.SkipVerify != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_SKIP_VERIFY=%s", context.SkipVerify))
		}

		if context.ClientTimeout != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_TIMEOUT=%s", context.ClientTimeout))
		}

		if context.ClusterAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLUSTER_ADDR=%s", context.ClusterAddr))
		}

		if context.License != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LICENSE=%s", context.License))
		}

		if context.LicensePath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LICENSE_PATH=%s", context.LicensePath))
		}

		if context.LogLevel != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LOG_LEVEL=%s", context.LogLevel))
		}

		if context.MaxRetries != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_MAX_RETRIES=%s", context.MaxRetries))
		}

		if context.RedirectAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_REDIRECT_ADDR=%s", context.RedirectAddr))
		}

		if context.TlsServerName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_TLS_SERVER_NAME=%s", context.TlsServerName))
		}

		if context.CliNoColour != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLI_NO_COLOR=%s", context.CliNoColour))
		}

		if context.RateLimit != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_RATE_LIMIT=%s", context.RateLimit))
		}

		if context.SvrLookup != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_SVR_LOOKUP=%s", context.SvrLookup))
		}

		if context.Mfa != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_MFA=%s", context.Mfa))
		}

		if context.HttpProxy != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_HTTP_PROXY=%s", context.HttpProxy))
		}

		if context.DisableRedirects != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_DISABLE_REDIRECTS=%s", context.DisableRedirects))
		}

		fmt.Println(strings.Join(exportCommands, "; "))
	},
}

var selectNomadCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target nomad select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Nomad[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Nomad[args[0]]

		exportCommands := []string{}

		if context.NomadEndpoint != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_ADDR=%s", context.NomadEndpoint))
		}

		if context.NomadToken != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_TOKEN=%s", context.NomadToken))
		}

		if context.NomadNamespace != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_NAMESPACE=%s", context.NomadNamespace))
		}

		if context.NomadCaCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CACERT=%s", context.NomadCaCert))
		}

		if context.NomadCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CLIENT_CERT=%s", context.NomadCert))
		}

		if context.NomadCaPath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CAPATH=%s", context.NomadCaPath))
		}

		if context.NomadKey != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CLIENT_KEY=%s", context.NomadKey))
		}

		if context.NomadRegion != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_REGION=%s", context.NomadRegion))
		}

		fmt.Println(strings.Join(exportCommands, "; "))
	},
}

var selectConsulCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target consul select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Consul[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Consul[args[0]]

		var exportCommandStr []string

		var shellCommandEndpoint string
		endpoint := context.ConsulEndpoint
		if endpoint != "" {
			shellCommandEndpoint = fmt.Sprintf("export CONSUL_HTTP_ADDR=%s", endpoint)
			exportCommandStr = append(exportCommandStr, shellCommandEndpoint)
		}

		var shellCommandToken string
		token := context.ConsulToken
		if token != "" {
			shellCommandToken = fmt.Sprintf("export CONSUL_HTTP_TOKEN=%s", token)
			exportCommandStr = append(exportCommandStr, shellCommandToken)
		}

		var shellCommandTokenFile string
		tokenFile := context.ConsulTokenFile
		if tokenFile != "" {
			shellCommandTokenFile = fmt.Sprintf("export CONSUL_HTTP_TOKEN_FILE=%s", tokenFile)
			exportCommandStr = append(exportCommandStr, shellCommandTokenFile)

		}

		var shellCommandCaCert string
		caCert := context.ConsulCaCert
		if caCert != "" {
			shellCommandCaCert = fmt.Sprintf("export CONSUL_CACERT=%s", caCert)
			exportCommandStr = append(exportCommandStr, shellCommandCaCert)
		}

		var shellCommandCert string
		cert := context.ConsulCert
		if cert != "" {
			shellCommandCert = fmt.Sprintf("export CONSUL_CLIENT_CERT=%s", cert)
			exportCommandStr = append(exportCommandStr, shellCommandCert)
		}

		var shellCommandCaPath string
		caPath := context.ConsulCaPath
		if caPath != "" {
			shellCommandCaPath = fmt.Sprintf("export CONSUL_CAPATH=%s", caPath)
			exportCommandStr = append(exportCommandStr, shellCommandCaPath)
		}

		var shellCommandKey string
		key := context.ConsulKey
		if key != "" {
			shellCommandKey = fmt.Sprintf("export CONSUL_CLIENT_KEY=%s", key)
			exportCommandStr = append(exportCommandStr, shellCommandKey)
		}

		var shellCommandNameSpace string
		namespace := context.ConsulNamespace
		if namespace != "" {
			shellCommandNameSpace = fmt.Sprintf("export CONSUL_NAMESPACE=%s", namespace)
			exportCommandStr = append(exportCommandStr, shellCommandNameSpace)
		}

		commandStr := strings.Join(exportCommandStr, "; ")
		fmt.Println(commandStr)
	},
}

var selectBoundaryCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target boundary select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Boundary[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Boundary[args[0]]

		var exportCommandStr []string

		var shellCommandEndpoint string
		endpoint := context.Endpoint
		if endpoint != "" {
			shellCommandEndpoint = fmt.Sprintf("export BOUNDARY_ADDR=%s", endpoint)
			exportCommandStr = append(exportCommandStr, shellCommandEndpoint)
		}

		var shellCommandToken string
		token := context.Token
		if token != "" {
			shellCommandToken = fmt.Sprintf("export NOMAD_TOKEN=%s", token) // TODO - Double check if this is needed
			exportCommandStr = append(exportCommandStr, shellCommandToken)
		}

		var shellCommandTokenName string
		tokenName := context.TokenName
		if tokenName != "" {
			shellCommandTokenName = fmt.Sprintf("export BOUNDARY_TOKEN_NAME=%s", tokenName)
			exportCommandStr = append(exportCommandStr, shellCommandTokenName)

		}

		var shellCommandCaCert string
		caCert := context.CaCert
		if caCert != "" {
			shellCommandCaCert = fmt.Sprintf("export BOUNDARY_CACERT=%s", caCert)
			exportCommandStr = append(exportCommandStr, shellCommandCaCert)
		}

		var shellCommandCert string
		cert := context.Cert
		if cert != "" {
			shellCommandCert = fmt.Sprintf("export BOUNDARY_CLIENT_CERT=%s", cert)
			exportCommandStr = append(exportCommandStr, shellCommandCert)
		}

		var shellCommandCaPath string
		caPath := context.CaPath
		if caPath != "" {
			shellCommandCaPath = fmt.Sprintf("export BOUNDARY_CAPATH=%s", caPath)
			exportCommandStr = append(exportCommandStr, shellCommandCaPath)
		}

		var shellCommandKey string
		key := context.Key
		if key != "" {
			shellCommandKey = fmt.Sprintf("export BOUNDARY_CLIENT_KEY=%s", key)
			exportCommandStr = append(exportCommandStr, shellCommandKey)
		}

		var shellCommandTlsInsecure string
		tlsInsecure := context.TlsInsecure
		if tlsInsecure != "" {
			shellCommandTlsInsecure = fmt.Sprintf("export BOUNDARY_TLS_INSECURE=%s", tlsInsecure)
			exportCommandStr = append(exportCommandStr, shellCommandTlsInsecure)
		}

		var shellCommandTlsServerName string
		tlsServerName := context.TlsServerName
		if tlsServerName != "" {
			shellCommandTlsServerName = fmt.Sprintf("export BOUNDARY_TLS_SERVER_NAME=%s", tlsServerName)
			exportCommandStr = append(exportCommandStr, shellCommandTlsServerName)
		}

		var shellCommandRecoveryConfig string
		recoveryConfig := context.RecoveryConfig
		if recoveryConfig != "" {
			shellCommandRecoveryConfig = fmt.Sprintf("export BOUNDARY_RECOVERY_CONFIG=%s", recoveryConfig)
			exportCommandStr = append(exportCommandStr, shellCommandRecoveryConfig)
		}

		var shellCommandConnectAuthZToken string
		connectAuthZToken := context.ConnectAuthZToken
		if connectAuthZToken != "" {
			shellCommandConnectAuthZToken = fmt.Sprintf("export BOUNDARY_CONNECT_AUTHZ_TOKEN=%s", connectAuthZToken)
			exportCommandStr = append(exportCommandStr, shellCommandConnectAuthZToken)
		}

		var shellCommandConnectExec string
		connectExec := context.ConnectExec
		if connectExec != "" {
			shellCommandConnectExec = fmt.Sprintf("export BOUNDARY_CONNECT_EXEC=%s", connectExec)
			exportCommandStr = append(exportCommandStr, shellCommandConnectExec)
		}

		var shellCommandConnectListenAddr string
		connectListenAddr := context.ConnectListenAddr
		if connectListenAddr != "" {
			shellCommandConnectListenAddr = fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_ADDR=%s", connectListenAddr)
			exportCommandStr = append(exportCommandStr, shellCommandConnectListenAddr)
		}

		var shellCommandConnectListenPort string
		connectListenPort := context.ConnectListenPort
		if connectListenPort != "" {
			shellCommandConnectListenPort = fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_PORT=%s", connectListenPort)
			exportCommandStr = append(exportCommandStr, shellCommandConnectListenPort)
		}

		var shellCommandConnectTargetScopeId string
		connectTargetScopeId := context.ConnectTargetScopeId
		if connectTargetScopeId != "" {
			shellCommandConnectTargetScopeId = fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_ID=%s", connectTargetScopeId)
			exportCommandStr = append(exportCommandStr, shellCommandConnectTargetScopeId)
		}

		var shellCommandConnectTargetScopeName string
		connectTargetScopeName := context.ConnectTargetScopeName
		if connectTargetScopeName != "" {
			shellCommandConnectTargetScopeName = fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_NAME=%s", connectTargetScopeName)
			exportCommandStr = append(exportCommandStr, shellCommandConnectTargetScopeName)
		}

		var shellCommandConnectAuthMethodId string
		authMethodId := context.AuthMethodId
		if authMethodId != "" {
			shellCommandConnectAuthMethodId = fmt.Sprintf("export BOUNDARY_AUTH_METHOD_ID=%s", authMethodId)
			exportCommandStr = append(exportCommandStr, shellCommandConnectAuthMethodId)
		}

		var shellCommandLogLevel string
		logLevel := context.LogLevel
		if logLevel != "" {
			shellCommandLogLevel = fmt.Sprintf("export BOUNDARY_LOG_LEVEL=%s", logLevel)
			exportCommandStr = append(exportCommandStr, shellCommandLogLevel)
		}

		var shellCommandFormat string
		format := context.Format
		if format != "" {
			shellCommandFormat = fmt.Sprintf("export BOUNDARY_CLI_FORMAT=%s", format)
			exportCommandStr = append(exportCommandStr, shellCommandFormat)
		}

		var shellCommandScopeId string
		scopeId := context.ScopeId
		if scopeId != "" {
			shellCommandScopeId = fmt.Sprintf("export BOUNDARY_SCOPE_ID=%s", scopeId)
			exportCommandStr = append(exportCommandStr, shellCommandScopeId)
		}

		commandStr := strings.Join(exportCommandStr, "; ")
		fmt.Println(commandStr)
	},
}

var selectTerraformCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target terraform select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Terraform[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Terraform[args[0]]

		var exportCommandStr []string

		for k, v := range context.Vars {
			command := fmt.Sprintf("export TF_VAR_%s=%s", k, v)
			exportCommandStr = append(exportCommandStr, command)
		}

		commandStr := strings.Join(exportCommandStr, "; ")
		fmt.Println(commandStr)
	},
}
