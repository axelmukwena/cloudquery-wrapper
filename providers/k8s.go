package providers

import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
)

// String to indicate provider folder
var kubernetesSubFolder string = "/.k8s"

// Create the credential file at provided location: "Users/username/.k8s/credential.json"
func setKubernetesCredentials(credentials string) {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}

	val := ValidateDir(homePath + kubernetesSubFolder)

	filename := homePath + kubernetesSubFolder + "/credentials"

	if val != nil {
		fmt.Println(val)
	}

	err := ioutil.WriteFile(filename, []byte(
		credentials,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to create Kubernetes credential file: %v\n", err)
	} else {
		fmt.Printf("Kubernetes credential file created at root\n")
	}
}

// Handle Kubernetes environment variables
func kubernetesEnvVariables(database string) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(homePath, err)
		return err
	}
	envVariables := "" +
		"export CQ_VAR_DSN=" + database + "\n" +
		"export KUBECONFIG=" + homePath + "/.k8s/credentials"
	val := CreateEnvFile(envVariables)

	return val
}

func Kubernetes(kubernetesString string, database string) (int, string) {
	success := 0
	message := ""
	setKubernetesCredentials(kubernetesString)
	val := kubernetesEnvVariables(database)

	if val != nil {
		fmt.Println(val)
		error := val.Error()
		return success, error
	}

	success, message = Fetch("k8s")

	return success, message
}
