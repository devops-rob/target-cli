package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "target",
	Short: "A context switcher CLI tool for Hashicorp Tools",
	Long: `Target allows a user to configure and switch between different contexts for their Vault, Nomad, Consul and Serf targets.
	
A context contains connection details for a given target.
Example: 
	a vault-dev context could point to 
	https://example-dev-vault.com:8200 with a vault token value is s.jidjibndiyuqepjepwo`,
	ValidArgs: []string{
		"vault",
		"nomad",
		"serf",
		"consul",
	},
	Args:    cobra.OnlyValidArgs,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(vaultCmd)
	rootCmd.AddCommand(nomadCmd)
	rootCmd.AddCommand(consulCmd)
	rootCmd.AddCommand(serfCmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".target" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".target")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
