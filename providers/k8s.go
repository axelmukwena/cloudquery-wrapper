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
func kubernetesEnvVariables() {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}
	envVariables := "export KUBECONFIG=" + homePath + "/.k8s/credentials"
	CreateEnvFile(envVariables)
}

func Kubernetes(kubernetesString string) int {
	success := 0
	setKubernetesCredentials(kubernetesString)
	kubernetesEnvVariables()
	success = Fetch("k8s")

	return success
}
