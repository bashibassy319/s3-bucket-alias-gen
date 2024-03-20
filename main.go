package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

type Config struct {
	AwsCredentialPath string
	BashrcPath        string
}

var homeDir = os.Getenv("HOME")

var defaultawsCredentialPath = homeDir + "/.aws/credentials"
var defaultbashrcPath = homeDir + "/.bashrc"

// Define command-line flags for configuration
var awsCredentialPath = flag.String("aws-credential-path", defaultawsCredentialPath, "Path to AWS credentials file, default path is "+defaultawsCredentialPath)
var bashrcPath = flag.String("bashrc-path", defaultbashrcPath, "Path to .bashrc file default path is "+defaultbashrcPath)
var dryRun = flag.Bool("dry-run", false, "Run in dry-run mode")
var help = flag.Bool("help", false, "Display help information") // Help flag

func main() {
	flag.Parse()

	if *help {
		flag.Usage() // Display usage information
		return
	}

	config := Config{
		AwsCredentialPath: *awsCredentialPath,
		BashrcPath:        *bashrcPath,
	}

	accessKey := promptAndValidate("ACCESS KEY", nil)
	secretKey := promptAndValidate("SECRET KEY", nil)
	aliasName := promptAndValidate("alias name", nil)
	endpoint := promptAndValidate("endpoint", nil)

	if *dryRun {
		fmt.Println("==================================")
		fmt.Println("In " + config.AwsCredentialPath)
		fmt.Printf("[%s]\n", aliasName)
		fmt.Printf("aws_access_key_id = %s\n", accessKey)
		fmt.Printf("aws_secret_access_key = %s\n", secretKey)

		fmt.Println("==================================")
		fmt.Println("In " + config.BashrcPath)
		fmt.Printf("alias %s='AWS_PROFILE=%s S3_ENDPOINT_URL=%s'\n", aliasName, aliasName, endpoint)
	} else {
		writeConfigToFile(config.AwsCredentialPath, accessKey, secretKey, aliasName, endpoint)
		writeBashrcToFile(config.BashrcPath, aliasName, endpoint)
	}
}

func promptAndValidate(label string, validate promptui.ValidateFunc) string {
	prompt := promptui.Prompt{
		Label:       label,
		Validate:    validate,
		HideEntered: true,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}
	return result
}

func writeConfigToFile(path, accessKey, secretKey, aliasName, endpoint string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "[%s]\n", aliasName)
	fmt.Fprintf(writer, "aws_access_key_id = %s\n", accessKey)
	fmt.Fprintf(writer, "aws_secret_access_key = %s\n", secretKey)
	writer.Flush()
}

func writeBashrcToFile(path, aliasName, endpoint string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "alias %s='AWS_PROFILE=%s S3_ENDPOINT_URL=%s'\n", aliasName, aliasName, endpoint)
	writer.Flush()
}

