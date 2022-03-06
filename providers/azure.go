package providers

import "C"
import "encoding/json"

// Azure JSON struct
type azureStruct struct {
	Azure_tenant_id       string
	Azure_client_id       string
	Azure_client_secret   string
	Azure_subscription_id string
}

// Parse and extract credentials
func parseAzure(azureString string) string {
	var azureStructNew azureStruct
	json.Unmarshal([]byte(azureString), &azureStructNew)

	envVariables := "" +
		"export AZURE_SUBSCRIPTION_ID=" + azureStructNew.Azure_subscription_id + "\n" +
		"export AZURE_TENANT_ID=" + azureStructNew.Azure_tenant_id + "\n" +
		"export AZURE_CLIENT_ID=" + azureStructNew.Azure_client_id + "\n" +
		"export AZURE_CLIENT_SECRET=" + azureStructNew.Azure_client_secret + "\n"

	return envVariables
}

func Azure(azureString string) int {
	envVariables := parseAzure(azureString)
	CreateEnvFile(envVariables)
	Fetch("azure")

	return 1
}
