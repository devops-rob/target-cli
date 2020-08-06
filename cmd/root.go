package cmd

import (
	"fmt"
	"os"
	"reflect"
	"target/pkg/targetdir"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var cfgFile = ""

var version string

// Config struct containing different product profiles
type Config struct {
	Vault   map[string]*Vault     `hcl:"vault,omitempty" mapstructure:"vault"`
	Consul  []map[string]*Consul  `hcl:"consul,omitempty" mapstructure:"consul"`
	Nomad   []map[string]*Nomad   `hcl:"nomad,omitempty" mapstructure:"nomad"`
	Serf    []map[string]*Serf    `hcl:"serf,omitempty" mapstructure:"serf"`
	Default []map[string]*Default `hcl:"default,omitempty" mapstructure:"default"`
}

// Vault struct with flag parameters
type Vault struct {
	Endpoint  string `hcl:"endpoint" mapstructure:"endpoint"`
	Token     string `hcl:"token,omitempty" mastructure:"token"`
	CaPath    string `hcl:"ca_path,omitempty" mapstructure:"ca_path"`
	CaCert    string `hcl:"ca_cert,omitempty" mapstructure:"ca_cert"`
	Cert      string `hcl:"cert,omitempty" mapstructure:"cert"`
	Key       string `hcl:"key,omitempty" mapstructure:"key"`
	Format    string `hcl:"format,omitempty" mapstructure:"format"`
	Namespace string `hcl:"namespace,omitempty" mapstructure:"namespace"`
}

var (
	c *Config
)

// Consul struct with flag parameters
type Consul struct {
	ConsulEndopoint string `hcl:"endpoint" mapstructure:"endpoint"`
	ConsulToken     string `hcl:"token,omitempty" mastructure:"token"`
	ConsulCaPath    string `hcl:"ca_path,omiempty" mapstructure:"ca_path"`
	ConsulCaCert    string `hcl:"ca_cert,omitempty" mapstructure:"ca_cert"`
	ConsulCert      string `hcl:"cert,omitempty" mapstructure:"cert"`
	ConsulKey       string `hcl:"key,omitempty" mapstructure:"key"`
	ConsulTokenFile string `hcl:"token_file,omitempty" mapstructure:"token_file"`
}

// Nomad struct with flag parameters
type Nomad struct {
	NomadEndopoint string `hcl:"endpoint" mapstructure:"endpoint"`
	NomadToken     string `hcl:"token,omitempty" mastructure:"token"`
	NomadCaPath    string `hcl:"ca_path,omiempty" mapstructure:"ca_path"`
	NomadCaCert    string `hcl:"ca_cert,omitempty" mapstructure:"ca_cert"`
	NomadCert      string `hcl:"cert,omitempty" mapstructure:"cert"`
	NomadKey       string `hcl:"key,omitempty" mapstructure:"key"`
	NomadRegion    string `hcl:"region,omitempty" mapstructure:"region"`
	NomadNamespace string `hcl:"namespace,omitempty" mapstructure:"namespace"`
}

// Serf struct with flag parameters
type Serf struct {
	SerfEndopoint string `hcl:"endpoint" mapstructure:"endpoint"`
	SerfToken     string `hcl:"token,omitempty" mastructure:"token"`
}

// Default struct with default profiles
type Default struct {
	VaultProfile  string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
	NomadProfile  string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
	ConsulProfile string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
	SerfProfile   string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
}

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
	targetdir.TargetHomeCreate()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(vaultCmd)
	rootCmd.AddCommand(nomadCmd)
	rootCmd.AddCommand(consulCmd)
	rootCmd.AddCommand(serfCmd)

}

// sliceOfMapsToMapHookFunc merges a slice of maps to a map
func sliceOfMapsToMapHookFunc() mapstructure.DecodeHookFunc {
	return func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		if from.Kind() == reflect.Slice && from.Elem().Kind() == reflect.Map && (to.Kind() == reflect.Struct || to.Kind() == reflect.Map) {
			source, ok := data.([]map[string]interface{})
			if !ok {
				return data, nil
			}
			if len(source) == 0 {
				return data, nil
			}
			if len(source) == 1 {
				return source[0], nil
			}
			// flatten the slice into one map
			convert := make(map[string]interface{})
			for _, mapItem := range source {
				for key, value := range mapItem {
					convert[key] = value
				}
			}
			return convert, nil
		}
		return data, nil
	}
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

		viper.AddConfigPath(home)
		viper.AddConfigPath("$HOME/.target")
		viper.SetConfigName("profiles")
		viper.SetConfigType("hcl")
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configOption := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
			sliceOfMapsToMapHookFunc(),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		))
		viper.Unmarshal(&c, configOption)
		if c.Vault == nil {
			c.Vault = map[string]*Vault{}
		}

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
