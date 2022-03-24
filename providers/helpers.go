package providers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Check if directory exists, if not, create it
func ValidateDir(dirName string) error {
	err := os.Mkdir(dirName, 0755)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory\n")
		}
		return nil
	}
	return err
}

// Create environment variable files
func CreateEnvFile(content string) {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}

	filename := homePath + "/.cloudquery-rails"

	err := ioutil.WriteFile(filename, []byte(
		content,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to create the environment variables file: %v\n", err)
	} else {
		fmt.Printf("Environment variables file created at root\n")
	}
}

// Fetch given provider
func Fetch(provider string) int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	config := "--config=" + currentDir + "/config/" + provider + ".hcl"

	cmd := exec.Command("cloudquery", "fetch", provider, config, "--enable-console-log")

	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)

	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}
