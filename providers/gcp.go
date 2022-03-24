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
func gcpEnvVariables() {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}
	envVariables := "export GOOGLE_APPLICATION_CREDENTIALS=" + homePath + "/.gcp/credentials.json"
	CreateEnvFile(envVariables)
}

func GCP(gcpString string) int {
	success := 0
	setGcpCredentials(gcpString)
	gcpEnvVariables()
	success = Fetch("gcp")

	return success
}
