package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
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

		// from here

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
		if mfa != "" {
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
		fmt.Println(commandStr)
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

		var exportCommandStr []string

		var shellCommandEndpoint string
		endpoint := context.NomadEndpoint
		if endpoint != "" {
			shellCommandEndpoint = fmt.Sprintf("export NOMAD_ADDR=%s", endpoint)
			exportCommandStr = append(exportCommandStr, shellCommandEndpoint)
		}

		var shellCommandToken string
		token := context.NomadToken
		if token != "" {
			shellCommandToken = fmt.Sprintf("export NOMAD_TOKEN=%s", token)
			exportCommandStr = append(exportCommandStr, shellCommandToken)
		}

		var shellCommandNamespace string
		namespace := context.NomadNamespace
		if namespace != "" {
			shellCommandNamespace = fmt.Sprintf("export NOMAD_NAMESPACE=%s", namespace)
			exportCommandStr = append(exportCommandStr, shellCommandNamespace)

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

		var shellCommandRegion string
		region := context.NomadRegion
		if region != "" {
			shellCommandRegion = fmt.Sprintf("export NOMAD_REGION=%s", region)
			exportCommandStr = append(exportCommandStr, shellCommandRegion)
		}

		commandStr := strings.Join(exportCommandStr, "; ")
		fmt.Println(commandStr)
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
