package providers

import "C"
import (
	"encoding/json"
	"fmt"
)

// Azure JSON struct
type azureStruct struct {
	Azure_tenant_id       string
	Azure_client_id       string
	Azure_client_secret   string
	Azure_subscription_id string
}

// Parse and extract credentials
func parseAzure(azureString string, database string) string {
	var azureStructNew azureStruct
	json.Unmarshal([]byte(azureString), &azureStructNew)

	envVariables := "" +
		"export CQ_VAR_DSN=" + database + "\n" +
		"export AZURE_SUBSCRIPTION_ID=" + azureStructNew.Azure_subscription_id + "\n" +
		"export AZURE_TENANT_ID=" + azureStructNew.Azure_tenant_id + "\n" +
		"export AZURE_CLIENT_ID=" + azureStructNew.Azure_client_id + "\n" +
		"export AZURE_CLIENT_SECRET=" + azureStructNew.Azure_client_secret + "\n"

	return envVariables
}

func Azure(azureString string, database string) (int, string) {
	success := 0
	message := ""
	envVariables := parseAzure(azureString, database)

	val := CreateEnvFile(envVariables)
	if val != nil {
		fmt.Println(val)
		error := val.Error()
		return success, error
	}

	success, message = Fetch("azure")

	return success, message
}
