package cmd

import (
	"github.com/spf13/cobra"
)

var (
	boundaryEndpoint               string
	boundaryToken                  string
	boundaryTokenName              string
	boundaryCaPath                 string
	boundaryCaCert                 string
	boundaryCert                   string
	boundaryKey                    string
	boundaryTlsInsecure            string
	boundaryTlsServerName          string
	boundaryRecoveryConfig         string
	boundaryConnectAuthZToken      string
	boundaryConnectExec            string
	boundaryConnectListenAddr      string
	boundaryConnectListenPort      string
	boundaryConnectTargetScopeId   string
	boundaryConnectTargetScopeName string
	boundaryAuthMethodId           string
	boundaryLogLevel               string
	boundaryCliFormat              string
	boundaryScopeId                string
)

var boundaryCmd = &cobra.Command{
	Use:     "boundary",
	Short:   "Manage Boundary context profiles ",
	Long:    `Manage Boundary context profiles.`,
	Example: `target boundary list`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"select",
		"update",
		"set-default",
	},
	Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
}

func init() {
	boundaryCmd.AddCommand(boundaryCreateCmd)
	boundaryCmd.AddCommand(deleteBoundaryCmd)
	boundaryCmd.AddCommand(selectBoundaryCmd)
	boundaryCmd.AddCommand(boundaryUpdateCmd)
	boundaryCmd.AddCommand(listBoundaryCmd)

}
