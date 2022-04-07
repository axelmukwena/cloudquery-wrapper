package main

// #include <stdlib.h>
import "C"
import (
	"cloudquery/providers"
)

// Main provider functions exported to Ruby

//export QueryAWS
func QueryAWS(awsString string, database string) *C.char {
	success, message, logfile := providers.AWS(awsString, database)

	// We read the contents later to allow time for cloudquery to save the log file
	output := providers.PrepareOutput(success, message, logfile)
	return C.CString(output)
}

//export QueryGCP
func QueryGCP(gcpString string, database string) *C.char {
	success, message, logfile := providers.GCP(gcpString, database)

	// We read the contents later to allow time for cloudquery to save the log file
	output := providers.PrepareOutput(success, message, logfile)
	return C.CString(output)
}

//export QueryAzure
func QueryAzure(azureString string, database string) *C.char {
	success, message, logfile := providers.Azure(azureString, database)

	// We read the contents later to allow time for cloudquery to save the log file
	output := providers.PrepareOutput(success, message, logfile)
	return C.CString(output)
}

//export QueryDigitalocean
func QueryDigitalocean(digitaloceanString string, database string) *C.char {
	success, message, logfile := providers.Digitalocean(digitaloceanString, database)

	// We read the contents later to allow time for cloudquery to save the log file
	output := providers.PrepareOutput(success, message, logfile)
	return C.CString(output)
}

//export QueryKubernetes
func QueryKubernetes(kubernetesString string, database string) *C.char {
	success, message, logfile := providers.Kubernetes(kubernetesString, database)

	// We read the contents later to allow time for cloudquery to save the log file
	output := providers.PrepareOutput(success, message, logfile)
	return C.CString(output)
}

//export QueryOkta
func QueryOkta(oktaString string, database string) *C.char {
	success, message, logfile := providers.Okta(oktaString, database)

	// We read the contents later to allow time for cloudquery to save the log file
	output := providers.PrepareOutput(success, message, logfile)
	return C.CString(output)
}

func main() {
}
