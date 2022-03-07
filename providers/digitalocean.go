package providers

import "C"
import "encoding/json"

// Digitalocean JSON struct
type digitaloceanStruct struct {
	Digitalocean_access_token string
	Digitalocean_token        string
	Spaces_access_key_id      string
	Spaces_secret_access_key  string
}

// Parse and extract credentials
func parseDigitalocean(digitaloceanString string) string {
	var digitaloceanStructNew digitaloceanStruct
	json.Unmarshal([]byte(digitaloceanString), &digitaloceanStructNew)

	envVariables := "" +
		"export DIGITALOCEAN_ACCESS_TOKEN=" + digitaloceanStructNew.Digitalocean_access_token + "\n" +
		"export DIGITALOCEAN_TOKEN=" + digitaloceanStructNew.Digitalocean_token + "\n" +
		"export SPACES_ACCESS_KEY_ID=" + digitaloceanStructNew.Spaces_access_key_id + "\n" +
		"export SPACES_SECRET_ACCESS_KEY=" + digitaloceanStructNew.Spaces_secret_access_key + "\n"

	return envVariables
}

func Digitalocean(digitaloceanString string) int {
	envVariables := parseDigitalocean(digitaloceanString)
	CreateEnvFile(envVariables)
	Fetch("digitalocean")

	return 1
}
