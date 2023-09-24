package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var boundaryUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `target boundary update example --endpoint="https://example-boundary.com:9200"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.Boundary[args[0]] == nil {
			log.Fatal("profile does not exists.  Try using the create command")

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
		err := viper.WriteConfig()
		if err != nil {
			return
		}

		fmt.Printf("Updated Boundary profile '%s'\n", args[0])
	},
}

func init() {
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-boundary.com:9200")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryToken, "token", "", "set boundary token")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryTokenName, "token-name", "", "Set boundary token name")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryTlsInsecure, "tls-insecure", "", "Set tls verification")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryTlsServerName, "tls-server-name", "", "Set the name of the server")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryRecoveryConfig, "recovery-config", "", "Set the boundary recovery config")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryConnectAuthZToken, "connect-authz-token", "", "Set the boundary token for connect authorisation")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryConnectExec, "connect-exec", "", "Set the binary that should be run in a connect session")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryConnectListenAddr, "connect-listen-addr", "", "Set the listen address for boundary connect")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryConnectListenPort, "connect-listen-port", "", "Set the listen port for boundary connect")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryConnectTargetScopeId, "connect-target-scope-id", "", "Set the scope id for the boundary connect target")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryConnectTargetScopeName, "connect-target-scope-name", "", "Set the scope name for the boundary connect target")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryAuthMethodId, "auth-method-id", "", "Set the boundary auth method id")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryLogLevel, "log-level", "", "Set the log level")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryCliFormat, "format", "", "Set the log output format")
	boundaryUpdateCmd.PersistentFlags().StringVar(&boundaryScopeId, "scope-id", "", "Set the boundary scope id to run commands in")

}
