package providers

import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
)

// String to indicate provider folder
var gcpSubFolder string = "/.gcp"

// Create the credential file at provided location: "Users/username/.gcp/credential.json"
func setGcpCredentials(credentials string) {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}

	val := ValidateDir(homePath + gcpSubFolder)

	filename := homePath + gcpSubFolder + "/credentials.json"

	if val != nil {
		fmt.Println(val)
	}

	err := ioutil.WriteFile(filename, []byte(
		credentials,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to create GCP credential file: %v\n", err)
	} else {
		fmt.Printf("GCP credential file created at root\n")
	}
}

// Handle GCP environment variables
func gcpEnvVariables(database string) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(homePath, err)
		return err
	}
	envVariables := "" +
		"export CQ_VAR_DSN=" + database + "\n" +
		"export GOOGLE_APPLICATION_CREDENTIALS=" + homePath + "/.gcp/credentials.json"
	val := CreateEnvFile(envVariables)

	return val
}

func GCP(gcpString string, database string) (bool, string, string) {
	success := false
	message := ""
	logs := ""
	setGcpCredentials(gcpString)
	val := gcpEnvVariables(database)

	if val != nil {
		fmt.Println(val)
		error := val.Error()
		return success, error, logs
	}

	success, message, logs = Fetch("gcp")

	return success, message, logs
}
