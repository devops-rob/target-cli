package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"
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
	Vault   map[string]*Vault   `json:"vault,omitempty" mapstructure:"vault"`
	Consul  map[string]*Consul  `json:"consul,omitempty" mapstructure:"consul"`
	Nomad   map[string]*Nomad   `json:"nomad,omitempty" mapstructure:"nomad"`
	Default map[string]*Default `json:"default,omitempty" mapstructure:"default"`
}

// Vault struct with flag parameters
type Vault struct {
	Endpoint  string `json:"endpoint,omitempty"`
	Token     string `json:"token,omitempty"`
	CaPath    string `json:"ca_path,omitempty"`
	CaCert    string `json:"ca_cert,omitempty"`
	Cert      string `json:"cert,omitempty"`
	Key       string `json:"key,omitempty"`
	Format    string `json:"format,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

var (
	c *Config
)

// Consul struct with flag parameters
type Consul struct {
	ConsulEndpoint  string `json:"endpoint" mapstructure:"endpoint"`
	ConsulToken     string `json:"token,omitempty" mastructure:"token"`
	ConsulCaPath    string `json:"ca_path,omiempty" mapstructure:"ca_path"`
	ConsulCaCert    string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	ConsulCert      string `json:"cert,omitempty" mapstructure:"cert"`
	ConsulKey       string `json:"key,omitempty" mapstructure:"key"`
	ConsulTokenFile string `json:"token_file,omitempty" mapstructure:"token_file"`
}

// Nomad struct with flag parameters
type Nomad struct {
	NomadEndpoint  string `json:"endpoint" mapstructure:"endpoint"`
	NomadToken     string `json:"token,omitempty" mastructure:"token"`
	NomadCaPath    string `json:"ca_path,omiempty" mapstructure:"ca_path"`
	NomadCaCert    string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	NomadCert      string `json:"cert,omitempty" mapstructure:"cert"`
	NomadKey       string `json:"key,omitempty" mapstructure:"key"`
	NomadRegion    string `json:"region,omitempty" mapstructure:"region"`
	NomadNamespace string `json:"namespace,omitempty" mapstructure:"namespace"`
}

// Default struct with default profiles
type Default struct {
	VaultProfile  string `json:"vault_profile,omitempty" mapstracture:"vault_profile"`
	NomadProfile  string `json:"vault_profile,omitempty" mapstracture:"vault_profile"`
	ConsulProfile string `json:"vault_profile,omitempty" mapstracture:"vault_profile"`
	SerfProfile   string `json:"vault_profile,omitempty" mapstracture:"vault_profile"`
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
	targetdir.TargetHomeCreate()

	rootCmd.AddCommand(vaultCmd)
	rootCmd.AddCommand(nomadCmd)
	rootCmd.AddCommand(consulCmd)

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

func initConfig() {
	var defaultConfig = &Config{
		// Your default configuration here...
	}

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath("$HOME/.target")
	viper.SetConfigName("profiles")
	viper.SetConfigType("json")

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, use default configuration
			c = defaultConfig
		} else {
			// Check if the error message indicates an empty file
			if strings.Contains(err.Error(), "empty") {
				// Config file is empty, use default configuration
				c = defaultConfig

				// You can set default values here if needed
				// For example:
				c.Vault = map[string]*Vault{
					"default": {
						Endpoint:  "https://example-vault-url.com",
						Token:     "your_vault_token",
						CaPath:    "",
						CaCert:    "",
						Cert:      "",
						Key:       "",
						Format:    "",
						Namespace: "",
					},
				}

			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	} else {
		// Config file found and successfully loaded
		configOption := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
			sliceOfMapsToMapHookFunc(),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		))
		if err := viper.Unmarshal(&c, configOption); err != nil {
			fmt.Println("Error unmarshaling config:", err)
			os.Exit(1)
		}

		if c.Vault == nil {
			c.Vault = map[string]*Vault{}
		}
		if c.Nomad == nil {
			c.Nomad = map[string]*Nomad{}
		}
		if c.Consul == nil {
			c.Consul = map[string]*Consul{}
		}
	}

	// Automatically bind environment variables
	viper.AutomaticEnv()

	// If a config file is found, print its path
	//if viper.ConfigFileUsed() != "" {
	//	fmt.Println("Using config file:", viper.ConfigFileUsed())
	//}
}
