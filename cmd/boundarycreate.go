package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var boundaryCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `target boundary create example --endpoint="https://example-boundary.com:9200"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Boundary[args[0]] != nil {
			log.Fatal("profile already exists")

		}
		b := &Boundary{
			Endpoint:               boundaryEndpoint,
			Token:                  boundaryToken,
			TokenName:              boundaryTokenName,
			CaPath:                 boundaryCaPath,
			CaCert:                 boundaryCaCert,
			Cert:                   boundaryCert,
			Key:                    boundaryKey,
			TlsInsecure:            boundaryTlsInsecure,
			TlsServerName:          boundaryTlsServerName,
			RecoveryConfig:         boundaryRecoveryConfig,
			ConnectAuthZToken:      boundaryConnectAuthZToken,
			ConnectExec:            boundaryConnectExec,
			ConnectListenAddr:      boundaryConnectListenAddr,
			ConnectListenPort:      boundaryConnectListenPort,
			ConnectTargetScopeId:   boundaryConnectTargetScopeId,
			ConnectTargetScopeName: boundaryConnectTargetScopeName,
			AuthMethodId:           boundaryAuthMethodId,
			LogLevel:               boundaryLogLevel,
			Format:                 boundaryCliFormat,
			ScopeId:                boundaryScopeId,
		}

		c.Boundary[args[0]] = b

		viper.Set("boundary", c.Boundary)
		viper.WriteConfig()

	},
}

func init() {

	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-boundary.com:9200")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryToken, "token", "", "set boundary token")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryTokenName, "token-name", "", "Set boundary token name")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryTlsInsecure, "tls-insecure", "", "Set tls verification")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryTlsServerName, "tls-server-name", "", "Set the name of the server")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryRecoveryConfig, "recovery-config", "", "Set the boundary recovery config")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryConnectAuthZToken, "connect-authz-token", "", "Set the boundary token for connect authorisation")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryConnectExec, "connect-exec", "", "Set the binary that should be run in a connect session")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryConnectListenAddr, "connect-listen-addr", "", "Set the listen address for boundary connect")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryConnectListenPort, "connect-listen-port", "", "Set the listen port for boundary connect")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryConnectTargetScopeId, "connect-target-scope-id", "", "Set the scope id for the boundary connect target")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryConnectTargetScopeName, "connect-target-scope-name", "", "Set the scope name for the boundary connect target")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryAuthMethodId, "auth-method-id", "", "Set the boundary auth method id")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryLogLevel, "log-level", "", "Set the log level")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryCliFormat, "format", "", "Set the log output format")
	boundaryCreateCmd.PersistentFlags().StringVar(&boundaryScopeId, "scope-id", "", "Set the boundary scope id to run commands in")

	boundaryCreateCmd.MarkPersistentFlagRequired(
		"endpoint",
	)
}
