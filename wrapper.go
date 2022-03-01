package main

import (
	"fmt"
	"io/ioutil"
)

func credentials() {
	err := ioutil.WriteFile("credentials", []byte(
		"[default]\n"+
			"aws_access_key_id=#####n"+
			"aws_secret_access_key=#####n",
	), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func config() {
	err := ioutil.WriteFile("config", []byte(
		"[default]\n"+
			"region=us-west-2\n"+
			"output=json\n",
	), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func main() {
	credentials()
	config()
}
