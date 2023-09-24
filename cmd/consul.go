package cmd

import (
	"errors"
	"fmt"
	"github.com/devops-rob/target-cli/pkg/targetdir"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var (
	consulEndpoint  string
	consulToken     string
	consulTokenFile string
	consulCaPath    string
	consulCaCert    string
	consulCert      string
	consulKey       string
	consulNamespace string
)

var consulCmd = &cobra.Command{
	Use:   "consul",
	Short: "Manage consul context profiles ",
	Long:  `Manage consul context profiles.`,
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

var consulSetDefaultCmd = &cobra.Command{
	Use:                   "set-default",
	Short:                 "set a default context profile for Consul ",
	Long:                  `set a default context profile for Consul.`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.Consul[args[0]] == nil {
			log.Fatalf("Profile %s not found", args[0])
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

		defaultScript := `
#!/bin/bash
` + commandStr

		absolutePath := targetdir.TargetHome() + "/defaults/consul.sh"

		file, err := os.OpenFile(absolutePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(defaultScript)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("default profile set")
	},
}

func init() {
	consulCmd.AddCommand(consulCreateCmd)
	consulCmd.AddCommand(consulSetDefaultCmd)
	consulCmd.AddCommand(selectConsulCmd)
	consulCmd.AddCommand(consulUpdateCmd)
	consulCmd.AddCommand(listConsulCmd)
	consulCmd.AddCommand(deleteConsulCmd)
}
