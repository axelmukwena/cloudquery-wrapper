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
	success, message := providers.AWS(awsString, database)

	successString := fmt.Sprint(success)
	output := "{\"success\":" + successString + ",\"message\":\"" + message + "\"}"
	return C.CString(output)
}

//export QueryGCP
func QueryGCP(gcpString string, database string) (int, string) {
	success, message := providers.GCP(gcpString, database)
	return success, message
}

//export QueryAzure
func QueryAzure(azureString string, database string) (int, string) {
	success, message := providers.Azure(azureString, database)
	return success, message
}

//export QueryDigitalocean
func QueryDigitalocean(digitaloceanString string, database string) (int, string) {
	success, message := providers.Digitalocean(digitaloceanString, database)
	return success, message
}

//export QueryKubernetes
func QueryKubernetes(kubernetesString string, database string) (int, string) {
	success, message := providers.Kubernetes(kubernetesString, database)
	return success, message
}

//export QueryOkta
func QueryOkta(oktaString string, database string) (int, string) {
	success, message := providers.Okta(oktaString, database)
	return success, message
}

func main() {
	// database := "tsdb://postgres:pass@localhost:5432/cloudtry?sslmode=disable"
	// credentials := "{\"aws_access_key_id\":\"AKIATE67RE6LE34ODKPN\",\"aws_secret_access_key\":\"lVUV9tdOO++aIX9IBDNXlAQfVDs5jGh2QNUlOo0\",\"region\":\"us-west-2\"}"
	// ifSuccess, message := providers.AWS(credentials, database)
	// fmt.Printf("Returned: %d", ifSuccess)
	// fmt.Printf("Message: %s", message)
}
