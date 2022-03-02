package main

import "C"

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Check if directory exists, if not, create it
func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, 0755)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory\n")
		}
		return nil
	}
	return err
}

// Create the credentials file at provided location: "Users/username/.aws/credential"
func SetCredentials(credentials string) {
	// String to indicate provider folder
	var subfolder string = "/.aaaa" // "/.aaaa" a temp placeholder for "/.aws"

	homepath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homepath, error)
	}

	val := ensureDir(homepath + subfolder)

	filename := homepath + subfolder + "/credentials"

	if val != nil {
		fmt.Println(val)
	}

	err := ioutil.WriteFile(filename, []byte(
		credentials,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to write file: %v\n", err)
	} else {
		fmt.Printf("Credential file created at root\n")
	}
}

// Create the config file at provided location: "Users/username/.aws/config"
func SetConfig(config string) {
	// String to indicate provider folder
	var subfolder string = "/.aaaa" // "/.aaaa" a temp placeholder for "/.aws"

	homepath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homepath, error)
	}

	val := ensureDir(homepath + subfolder)

	filename := homepath + subfolder + "/config"

	if val != nil {
		fmt.Println(val)
	}

	err := ioutil.WriteFile(filename, []byte(
		config,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to write file: %v\n", err)
	} else {
		fmt.Printf("Config file created at root\n")
	}
}

func cloudquery() {
	cmd := exec.Command("cloudquery", "fetch")

	// err := cmd.Run()
	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}

//export QueryAWS
func QueryAWS(credentials, config string) int {
	// Main AWS funtion exported to Ruby

	SetCredentials(credentials)
	SetConfig(config)
	cloudquery()

	return 1 // 0 if fail. Easier to send int than boolean
}

func main() {}
