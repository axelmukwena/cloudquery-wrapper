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
func CreateEnvFile(content string) error {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}

	filename := homePath + "/.krishancy-cloudquery"

	err := ioutil.WriteFile(filename, []byte(
		content,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to create the environment variables file: %v\n", err)
		return err
	} else {
		fmt.Printf("Environment variables file created at root\n")
	}
	return nil
}

// Fetch given provider
func Fetch(provider string) (bool, string, string) {
	logs := ""
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		error := err.Error()
		return false, error, logs
	}

	// Uncommnent below for Ruby project
	// config := "--config=" + currentDir + "/config/initializers/cloudquery/config/" + provider + ".hcl"
	config := "--config=" + currentDir + "/config/" + provider + ".hcl"

	cmd := exec.Command("cloudquery", "fetch", provider, config, "--enable-console-log")

	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)

	// sep := []byte("\n")
	logs = string(stdoutStderr)

	if err != nil {
		log.Println(err)
		error := err.Error()
		return false, error, logs
	}
	return true, "success", logs
}
