package providers

import "C"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// String to indicate provider folder
var awsSubFolder string = "/.aws" // "/.aaaa" a temp placeholder for "/.aws"

// AWS JSON struct
type awsStruct struct {
	Aws_access_key_id     string
	Aws_secret_access_key string
	Aws_session_token     string
	Region                string
}

func parseAWS(awsString string) (string, string) {
	var awsStructNew awsStruct
	json.Unmarshal([]byte(awsString), &awsStructNew)

	// Extract credentials
	credentials := string("")
	if awsStructNew.Aws_session_token != "" {
		credentials = "[temp]\n" +
			"aws_access_key_id = " + awsStructNew.Aws_access_key_id + "\n" +
			"aws_secret_access_key = " + awsStructNew.Aws_secret_access_key + "\n" +
			"aws_session_token = " + awsStructNew.Aws_session_token + "\n"
	} else {
		credentials = "[default]\n" +
			"aws_access_key_id = " + awsStructNew.Aws_access_key_id + "\n" +
			"aws_secret_access_key = " + awsStructNew.Aws_secret_access_key + "\n"
	}

	// Extract config
	config := string("")
	if awsStructNew.Region != "" {
		config = "[default]\nregion = " + awsStructNew.Region + "\n"
	} else {
		config = "[default]\nregion = us-west-2\n"
	}

	return credentials, config
}

// Create the credentials file at provided location: "Users/username/.aws/credential"
func setAwsCredentials(credentials string) {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}

	val := ValidateDir(homePath + awsSubFolder)

	filename := homePath + awsSubFolder + "/credentials"

	if val != nil {
		fmt.Println(val)
	}

	err := ioutil.WriteFile(filename, []byte(
		credentials,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to create AWS credential file: %v\n", err)
	} else {
		fmt.Printf("AWS credential file created at root\n")
	}
}

// Create the config file at provided location: "Users/username/.aws/config"
func setAwsConfig(config string) {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}

	val := ValidateDir(homePath + awsSubFolder)

	filename := homePath + awsSubFolder + "/config"

	if val != nil {
		fmt.Println(val)
	}

	err := ioutil.WriteFile(filename, []byte(
		config,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to create AWS config file: %v\n", err)
	} else {
		fmt.Printf("AWS config file created at root\n")
	}
}

// Parse and extract credentials
func parseAws(database string) string {
	envVariables := string("") +
		"export CQ_VAR_DSN=" + database + "\n"

	return envVariables
}

func AWS(awsString string, database string) (bool, string, string) {
	success := false
	message := string("")
	logfile := string("")
	credentials, config := parseAWS(awsString)

	envVariables := parseAws(database)
	val := CreateEnvFile(envVariables)
	if val != nil {
		fmt.Println(val)
		error := val.Error()
		return success, error, ""
	}

	setAwsCredentials(credentials)
	setAwsConfig(config)
	success, message, logfile = Fetch("aws")

	return success, message, logfile
}
