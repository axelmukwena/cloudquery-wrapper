package providers

import "C"
import (
	"encoding/json"
	"fmt"
)

// Okta JSON struct
type oktaStruct struct {
	Okta_domain    string
	Okta_api_token string
}

// Parse and extract credentials
func parseOkta(oktaString string, database string) string {
	var oktaStructNew oktaStruct
	json.Unmarshal([]byte(oktaString), &oktaStructNew)

	envVariables := "" +
		"export CQ_VAR_DSN=" + database + "\n" +
		"export CQ_VAR_OKTA_DOMAIN=" + oktaStructNew.Okta_domain + "\n" +
		"export OKTA_API_TOKEN=" + oktaStructNew.Okta_api_token + "\n"

	return envVariables
}

func Okta(oktaString string, database string) (int, string) {
	success := 0
	message := ""
	envVariables := parseOkta(oktaString, database)
	val := CreateEnvFile(envVariables)

	if val != nil {
		fmt.Println(val)
		error := val.Error()
		return success, error
	}

	success, message = Fetch("okta")

	return success, message
}
