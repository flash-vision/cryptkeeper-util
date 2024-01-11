package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/flash-vision/cryptkeeper" // Replace with your actual module path
)

func main() {
	// Define command-line flags
	passkey := flag.String("passkey", "", "Passkey for encryption/decryption")
	configFile := flag.String("configfile", "", "Path to the config file")
	envVar := flag.String("env-var", "CRYPTKEEPER_CRYPTKEY", "Environment variable name for the passkey")

	// Parse the flags
	flag.Parse()

	// Check if required flags are provided
	if *passkey == "" || *configFile == "" {
		log.Fatal("Both flags -passkey and -configfile are required")
	}

	// Set the environment variable for the passkey
	os.Setenv(*envVar, *passkey)

	// Create a ConfigFile instance
	cf := cryptkeeper.NewConfigFile(*configFile, *envVar)

	// Perform encryption/decryption based on the file extension
	if err := cf.ProcessFile(); err != nil {
		log.Fatalf("Failed to process file: %v", err)
	}

	fmt.Printf("File successfully processed. Output at: %s\n", cf.FilePath)
}
