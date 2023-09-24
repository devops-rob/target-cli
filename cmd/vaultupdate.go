package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// updateCmd represents the update command
var vaultUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target vault update example --endpoint="https://example2-vault.com:8200" --token="t.loejwikdjuidfhjdi"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Vault[args[0]] == nil {
			log.Fatal("profile does not exists.  Try using the create command")

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

		fmt.Printf("Updated Vault profile %s\n", args[0])

	},
}

func init() {
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-vault.com:8200")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultToken, "token", "", "set vault auth token for this context")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultFormat, "format", "", `set vault output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultNamespace, "namespace", "", "set vault namespace to use for command")

	vaultUpdateCmd.PersistentFlags().StringVar(&vaultSkipVerify, "skip-verify", "", "Do not verify Vault's presented certificate before communicating with it")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultClientTimeout, "client-timeout", "", "Set the Timeout variable")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultClusterAddr, "cluster-addr", "", "Set the address that should be used for other cluster members to connect to this node when in High Availability mode")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultLicense, "license", "", "Specify a license to use for this node.")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultLicensePath, "license-path", "", "Specify a path to a license on disk to use for this node.")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultLogLevel, "log-level", "", "Set the Vault server's log level")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultMaxRetries, "max-retries", "", "Set the maximum number of retries when certain error codes are encountered")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultRedirectAddr, "redirect-addr", "", "Set the address that should be used when clients are redirected to this node when in High Availability mode")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultTlsServerName, "tls-server-name", "", "Set the name to use as the SNI host when connecting via TLS")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultCliNoColour, "cli-no-colour", "", "If provided, Vault output will not include ANSI color escape sequence characters")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultRateLimit, "rate-limit", "", "Set the rate at which the vault command sends requests to Vault")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultSvrLookup, "svr-lookup", "", "Enables the client to lookup the host through DNS SRV look up")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultMfa, "mfa", "", "Set the MFA credentials in the format mfa_method_name[:key[=value]]")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultHttpProxy, "http=proxy", "", "Set the HTTP or HTTPS proxy location which should be used by all requests to access Vault")
	vaultUpdateCmd.PersistentFlags().StringVar(&vaultDisableRedirects, "disable-redirects", "", "Prevents the Vault client from following redirects")

}
