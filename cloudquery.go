package main

// #include <stdlib.h>
import "C"
import (
	"cloudquery/providers"
	"fmt"
)

// Main provider functions exported to Ruby

//export QueryAWS
func QueryAWS(awsString string, database string) *C.char {
	success, message, logs := providers.AWS(awsString, database)

	successString := fmt.Sprint(success)
	output := string("{\"success\":" + successString + ", \"message\":" + message + ", \"logs\": " + logs + "}")
	return C.CString(output)
}

//export QueryGCP
func QueryGCP(gcpString string, database string) *C.char {
	success, message, logs := providers.GCP(gcpString, database)

	successString := fmt.Sprint(success)
	output := string("{\"success\":" + successString + ", \"message\":" + message + ", \"logs\": " + logs + "}")
	return C.CString(output)
}

//export QueryAzure
func QueryAzure(azureString string, database string) *C.char {
	success, message, logs := providers.Azure(azureString, database)

	successString := fmt.Sprint(success)
	output := string("{\"success\":" + successString + ", \"message\":" + message + ", \"logs\": " + logs + "}")
	return C.CString(output)
}

//export QueryDigitalocean
func QueryDigitalocean(digitaloceanString string, database string) *C.char {
	success, message, logs := providers.Digitalocean(digitaloceanString, database)

	successString := fmt.Sprint(success)
	output := string("{\"success\":" + successString + ", \"message\":" + message + ", \"logs\": " + logs + "}")
	return C.CString(output)
}

//export QueryKubernetes
func QueryKubernetes(kubernetesString string, database string) *C.char {
	success, message, logs := providers.Kubernetes(kubernetesString, database)

	successString := fmt.Sprint(success)
	output := string("{\"success\":" + successString + ", \"message\":" + message + ", \"logs\": " + logs + "}")
	return C.CString(output)
}

//export QueryOkta
func QueryOkta(oktaString string, database string) *C.char {
	success, message, logs := providers.Okta(oktaString, database)

	successString := fmt.Sprint(success)
	output := string("{\"success\":" + successString + ", \"message\":" + message + ", \"logs\": " + logs + "}")
	return C.CString(output)
}

func main() {
}
