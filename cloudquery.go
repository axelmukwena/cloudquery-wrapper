package main

import "C"
import (
	"cloudquery/providers"
)

// Main provider functions exported to Ruby

//export QueryAWS
func QueryAWS(awsString string) int {
	ifSuccess := providers.AWS(awsString)
	return ifSuccess
}

//export QueryGCP
func QueryGCP(gcpString string) int {
	ifSuccess := providers.GCP(gcpString)
	return ifSuccess
}

//export QueryAzure
func QueryAzure(azureString string) int {
	ifSuccess := providers.Azure(azureString)
	return ifSuccess
}

//export QueryDigitalocean
func QueryDigitalocean(digitaloceanString string) int {
	ifSuccess := providers.Digitalocean(digitaloceanString)
	return ifSuccess
}

func main() {
	// ifSuccess := providers.GCP("myString")
	// fmt.Printf("Returned: %v", ifSuccess)
}
