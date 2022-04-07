package providers

import "C"
import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

// Check if directory exists, if not, create it
func ValidateDir(dirName string) error {
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

// Create environment variable files
func CreateEnvFile(content string) error {
	homePath, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(homePath, error)
	}

	filename := homePath + "/.krishancy-cloudquery"

	err := ioutil.WriteFile(filename, []byte(
		content,
	), 0777)
	if err != nil {
		fmt.Printf("Unable to create the environment variables file: %v\n", err)
		return err
	} else {
		fmt.Printf("Environment variables file created at root\n")
	}
	return nil
}

func Configurations(provider string) (bool, string, string, string, string, string) {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		error := err.Error()
		return false, error, string(""), string(""), string(""), string("")
	}

	// Toggle comments below for Ruby project

	// Create log folder if it doesn't exist
	logsDir := "/config/initializers/cloudquery/logs"
	// logsDir := "/config/logs"
	val := ValidateDir(currentDir + logsDir)
	if val != nil {
		fmt.Println(val)
	}
	logsDirConfig := "--log-directory=" + currentDir + logsDir

	filename := fmt.Sprint(time.Now().Unix())
	logFileConfig := "--log-file=" + filename + ".log"

	config := "--config=" + currentDir + "/config/initializers/cloudquery/config/" + provider + ".hcl"
	//config := "--config=" + currentDir + "/config/" + provider + ".hcl"

	return true, string("success"), config, logsDirConfig, logFileConfig, filename
}

// Fetch given provider
func Fetch(provider string) (bool, string, string) {

	success, message, config, logsDirConfig, logFileConfig, filename := Configurations(provider)

	if success == false {
		return false, message, string("")
	}

	cmd := exec.Command("cloudquery", "fetch", provider, config, "--enable-console-log", logsDirConfig, logFileConfig)

	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)

	// sep := []byte("\n")
	// logs = fmt.Sprint(stdoutStderr)
	// logs = string(stdoutStderr)

	if err != nil {
		log.Println(err)
		error := err.Error()
		return false, error, filename
	}
	return success, message, filename
}

// Get the latest config file
func ReadLogFile(filename string) string {

	currentDir, err := os.Getwd()

	if err != nil {
		log.Println(err)
		// error := err.Error()
		// return error, ""
		return "[]"
	}

	logsDir := currentDir + "/config/initializers/cloudquery/logs/"
	// logsDir := "/logs"

	logFile := logsDir + filename + ".log"

	file, err := os.Open(logFile)
	if err != nil {
		log.Println(err)
		return "[]"
	}
	defer file.Close()

	contents := string("")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if contents == "" {
			contents += scanner.Text()
		} else {
			contents += "," + scanner.Text()
		}

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return "[]"
	}

	return "[" + contents + "]"
}

func PrepareOutput(success bool, message string, logfile string) string {

	logs := string("[]")
	logs = ReadLogFile(logfile)

	successString := fmt.Sprint(success)
	output := string("{\"success\":" + successString + ", \"message\":\"" + message + "\", \"logfile\":\"" + logfile + "\", \"logs\":" + logs + "}")
	return output
}
