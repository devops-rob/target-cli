package cmd

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

type ConfigurationFlags struct {
	path string `json:"path,omitempty" mapstructure:"path"`
}

var path string

var configlCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure target CLI for shell sessions",
	Long:  `Configure target CLI for shell sessions.`,
	//ValidArgs: []string{
	//	"zsh",
	//	"bash",
	//	"fish",
	//	"dash",
	//	"ksh",
	//	"powershell",
	//},
	Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}

		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}

		//defer file.Close()

		scanner := bufio.NewScanner(file)

		searchText := "# Target CLI Defaults"

		// Loop through each line of the file
		lineNumber := 0
		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()
			// Check if the line contains the specific text
			if strings.Contains(line, searchText) {
				log.Fatalf("Target CLI Configuration already written to this file")
			}
		}

		// Check for scanner error
		if err := scanner.Err(); err != nil {
			log.Fatalf("scanner error: %s", err)
		}

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fullTargetPath := home + "/.target/defaults/*"

		configScript := `
# Target CLI Defaults

for file in ` + fullTargetPath + `; do
    if [ -f "$file" ]; then
        source "$file"
    fi
done
`
		file2, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//defer file2.Close()
		_, err = file2.WriteString(configScript)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	configlCmd.PersistentFlags().StringVar(&path, "path", "", "Absolute path for your chosen shell configuration file")

	configlCmd.MarkPersistentFlagRequired("path")
}
