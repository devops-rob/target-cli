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

		exportCommandStr := []string{}

		if context.ConsulEndpoint != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_HTTP_ADDR=%s", context.ConsulEndpoint))
		}

		if context.ConsulToken != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_HTTP_TOKEN=%s", context.ConsulToken))
		}

		if context.ConsulTokenFile != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_HTTP_TOKEN_FILE=%s", context.ConsulTokenFile))
		}

		if context.ConsulCaCert != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CACERT=%s", context.ConsulCaCert))
		}

		if context.ConsulCert != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CLIENT_CERT=%s", context.ConsulCert))
		}

		if context.ConsulCaPath != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CAPATH=%s", context.ConsulCaPath))
		}

		if context.ConsulKey != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CLIENT_KEY=%s", context.ConsulKey))
		}

		if context.ConsulNamespace != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_NAMESPACE=%s", context.ConsulNamespace))
		}

		fmt.Println(strings.Join(exportCommandStr, "; "))
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

		exportCommands := []string{}

		if context.Endpoint != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_ADDR=%s", context.Endpoint))
		}

		// NOTE: users can use either a "BOUNARY_TOKEN" (for a file on disk) or
		// a "BOUNDARY_TOKEN_NAME" (for platform-specific OS credential store).

		if context.Token != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TOKEN=%s", context.Token))
		}

		if context.TokenName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TOKEN_NAME=%s", context.TokenName))
		}

		if context.CaCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CACERT=%s", context.CaCert))
		}

		if context.Cert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLIENT_CERT=%s", context.Cert))
		}

		if context.CaPath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CAPATH=%s", context.CaPath))
		}

		if context.Key != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLIENT_KEY=%s", context.Key))
		}

		if context.TlsInsecure != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TLS_INSECURE=%s", context.TlsInsecure))
		}

		if context.TlsServerName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TLS_SERVER_NAME=%s", context.TlsServerName))
		}

		if context.RecoveryConfig != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_RECOVERY_CONFIG=%s", context.RecoveryConfig))
		}

		if context.ConnectAuthZToken != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_AUTHZ_TOKEN=%s", context.ConnectAuthZToken))
		}

		if context.ConnectExec != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_EXEC=%s", context.ConnectExec))
		}

		if context.ConnectListenAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_ADDR=%s", context.ConnectListenAddr))
		}

		if context.ConnectListenPort != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_PORT=%s", context.ConnectListenPort))
		}

		if context.ConnectTargetScopeId != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_ID=%s", context.ConnectTargetScopeId))
		}

		if context.ConnectTargetScopeName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_NAME=%s", context.ConnectTargetScopeName))
		}

		if context.AuthMethodId != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_AUTH_METHOD_ID=%s", context.AuthMethodId))
		}

		if context.LogLevel != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_LOG_LEVEL=%s", context.LogLevel))
		}

		if context.Format != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLI_FORMAT=%s", context.Format))
		}

		if context.ScopeId != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_SCOPE_ID=%s", context.ScopeId))
		}

		fmt.Println(strings.Join(exportCommands, "; "))
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
