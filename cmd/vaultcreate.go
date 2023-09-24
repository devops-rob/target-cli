package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var vaultCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target vault create example --endpoint="https://example-vault.com:8200" --token="s.giqoewbnmdjalkjk"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.Vault[args[0]] != nil {
			log.Fatal("profile already exists")

		}
		v := &Vault{
			Endpoint:         vaultEndpoint,
			Token:            vaultToken,
			CaPath:           vaultCaPath,
			CaCert:           vaultCaCert,
			Cert:             vaultCert,
			Key:              vaultKey,
			Format:           vaultFormat,
			Namespace:        vaultNamespace,
			SkipVerify:       vaultSkipVerify,
			ClientTimeout:    vaultClientTimeout,
			ClusterAddr:      vaultClusterAddr,
			License:          vaultLicense,
			LicensePath:      vaultLicensePath,
			LogLevel:         vaultLogLevel,
			MaxRetries:       vaultMaxRetries,
			RedirectAddr:     vaultRedirectAddr,
			TlsServerName:    vaultTlsServerName,
			CliNoColour:      vaultCliNoColour,
			RateLimit:        vaultRateLimit,
			SvrLookup:        vaultSvrLookup,
			Mfa:              vaultMfa,
			HttpProxy:        vaultHttpProxy,
			DisableRedirects: vaultDisableRedirects,
		}

		c.Vault[args[0]] = v

		viper.Set("vault", c.Vault)
		err := viper.WriteConfig()
		if err != nil {
			return
		}
		fmt.Printf("Created Vault profile '%s'\n", args[0])

	},
}

func init() {

	vaultCreateCmd.PersistentFlags().StringVar(&vaultEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultToken, "token", "", "set vault auth token for this context")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultFormat, "format", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultCreateCmd.PersistentFlags().StringVar(&vaultNamespace, "namespace", "", "set vault namespace to use for command")
	// new flags
	vaultCreateCmd.PersistentFlags().StringVar(&vaultSkipVerify, "skip-verify", "", "Do not verify Vault's presented certificate before communicating with it")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultClientTimeout, "client-timeout", "", "Set the Timeout variable")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultClusterAddr, "cluster-addr", "", "Set the address that should be used for other cluster members to connect to this node when in High Availability mode")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultLicense, "license", "", "Specify a license to use for this node.")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultLicensePath, "license-path", "", "Specify a path to a license on disk to use for this node.")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultLogLevel, "log-level", "", "Set the Vault server's log level")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultMaxRetries, "max-retries", "", "Set the maximum number of retries when certain error codes are encountered")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultRedirectAddr, "redirect-addr", "", "Set the address that should be used when clients are redirected to this node when in High Availability mode")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultTlsServerName, "tls-server-name", "", "Set the name to use as the SNI host when connecting via TLS")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultCliNoColour, "cli-no-colour", "", "If provided, Vault output will not include ANSI color escape sequence characters")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultRateLimit, "rate-limit", "", "Set the rate at which the vault command sends requests to Vault")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultSvrLookup, "svr-lookup", "", "Enables the client to lookup the host through DNS SRV look up")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultMfa, "mfa", "", "Set the MFA credentials in the format mfa_method_name[:key[=value]]")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultHttpProxy, "http=proxy", "", "Set the HTTP or HTTPS proxy location which should be used by all requests to access Vault")
	vaultCreateCmd.PersistentFlags().StringVar(&vaultDisableRedirects, "disable-redirects", "", "Prevents the Vault client from following redirects")

	vaultCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)

}
