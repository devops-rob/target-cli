package targetdir

import (
	"fmt"
	"os"
	"runtime"
)

// HomeFolder returns the users homefolder this will be $HOME on windows and mac and
// USERPROFILE on windows
func HomeFolder() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")

	}

	return os.Getenv("HOME")
}

// TargetHome returns the location of the target
// folder, usually $HOME/.target
func TargetHome() string {
	return fmt.Sprintf("%s/.target", HomeFolder())
}

// TargetHomeCreate checks for the target directory
// and profiles.json file and creates if they don't exist
func TargetHomeCreate() {
	var defaultConfig = "{\n\t\"vault\": {},\n\t\"consul\": {},\n\t\"nomad\": {},\n\t\"terraform\": {}\n}"
	targetHome := TargetHome()
	if _, err := os.Stat(targetHome); os.IsNotExist(err) {
		os.Mkdir(targetHome, 0755)
	}

	f := fmt.Sprintf("%s/profiles.json", targetHome)

	if _, err := os.Stat(f); os.IsNotExist(err) {
		// Create and write the default configuration to the file
		err := os.WriteFile(f, []byte(defaultConfig), 0644)
		if err != nil {
			fmt.Printf("Error creating and writing to profiles.json: %v\n", err)
		}
	}

	defaultsDir := targetHome + "/defaults"
	if _, err := os.Stat(defaultsDir); os.IsNotExist(err) {
		os.Mkdir(defaultsDir, 0755)
	}
}
