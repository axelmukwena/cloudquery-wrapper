package providers

import "C"
import "encoding/json"

// Okta JSON struct
type oktaStruct struct {
	Okta_domain    string
	Okta_api_token string
}

// Parse and extract credentials
func parseOkta(oktaString string) string {
	var oktaStructNew oktaStruct
	json.Unmarshal([]byte(oktaString), &oktaStructNew)

	envVariables := "" +
		"export CQ_VAR_OKTA_DOMAIN=" + oktaStructNew.Okta_domain + "\n" +
		"export OKTA_API_TOKEN=" + oktaStructNew.Okta_api_token + "\n"

	return envVariables
}

func Okta(oktaString string) int {
	success := 0
	envVariables := parseOkta(oktaString)
	CreateEnvFile(envVariables)
	success = Fetch("okta")

	return success
}
