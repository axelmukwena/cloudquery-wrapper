package main

import "C"
import (
	"cloudquery/providers"
)

// Main provider functions exported to Ruby

//export QueryAWS
func QueryAWS(awsString string) int {
	success := providers.AWS(awsString)
	return success
}

//export QueryGCP
func QueryGCP(gcpString string) int {
	success := providers.GCP(gcpString)
	return success
}

//export QueryAzure
func QueryAzure(azureString string) int {
	success := providers.Azure(azureString)
	return success
}

//export QueryDigitalocean
func QueryDigitalocean(digitaloceanString string) int {
	success := providers.Digitalocean(digitaloceanString)
	return success
}

//export QueryKubernetes
func QueryKubernetes(kubernetesString string) int {
	success := providers.Kubernetes(kubernetesString)
	return success
}

//export QueryOkta
func QueryOkta(oktaString string) int {
	success := providers.Okta(oktaString)
	return success
}

func main() {
	// ifSuccess := providers.GCP("myString")
	// fmt.Printf("Returned: %v", ifSuccess)
}
