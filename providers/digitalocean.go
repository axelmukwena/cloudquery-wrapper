package providers

import "C"
import (
	"encoding/json"
	"fmt"
)

// Digitalocean JSON struct
type digitaloceanStruct struct {
	Digitalocean_access_token string
	Digitalocean_token        string
	Spaces_access_key_id      string
	Spaces_secret_access_key  string
}

// Parse and extract credentials
func parseDigitalocean(digitaloceanString string, database string) string {
	var digitaloceanStructNew digitaloceanStruct
	json.Unmarshal([]byte(digitaloceanString), &digitaloceanStructNew)

	envVariables := "" +
		"export CQ_VAR_DSN=" + database + "\n" +
		"export DIGITALOCEAN_ACCESS_TOKEN=" + digitaloceanStructNew.Digitalocean_access_token + "\n" +
		"export DIGITALOCEAN_TOKEN=" + digitaloceanStructNew.Digitalocean_token + "\n" +
		"export SPACES_ACCESS_KEY_ID=" + digitaloceanStructNew.Spaces_access_key_id + "\n" +
		"export SPACES_SECRET_ACCESS_KEY=" + digitaloceanStructNew.Spaces_secret_access_key + "\n"

	return envVariables
}

func Digitalocean(digitaloceanString string, database string) (bool, string, string) {
	success := false
	message := ""
	logs := ""
	envVariables := parseDigitalocean(digitaloceanString, database)

	val := CreateEnvFile(envVariables)
	if val != nil {
		fmt.Println(val)
		error := val.Error()
		return success, error, logs
	}

	success, message, logs = Fetch("digitalocean")

	return success, message, logs
}
