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
	mode := flag.String("mode", "", "Mode of operation: encrypt or decrypt")
	passkey := flag.String("passkey", "", "Passkey for encryption/decryption")
	configFile := flag.String("configfile", "", "Path to the config file")

	// Parse the flags
	flag.Parse()

	// Check if required flags are provided
	if *mode == "" || *passkey == "" || *configFile == "" {
		log.Fatal("All flags -mode, -passkey, and -configfile are required")
	}

	// Set the environment variable for the passkey
	os.Setenv("CRYPTKEEPER_CRYPTKEY", *passkey)

	var destinationFile string
	if *mode == "encrypt" {
		destinationFile = *configFile + ".encrypted"
	} else if *mode == "decrypt" {
		destinationFile = *configFile
		*configFile += ".encrypted"
	} else {
		log.Fatal("Invalid mode: choose 'encrypt' or 'decrypt'")
	}

	// Create a ConfigFile instance
	cf := cryptkeeper.NewConfigFile(*configFile, *mode == "encrypt", destinationFile)

	// Perform encryption/decryption
	if err := cf.ProcessFile(); err != nil {
		log.Fatalf("Failed to process file: %v", err)
	}

	// Remove the original file after successful encryption/decryption
	if err := os.Remove(*configFile); err != nil {
		log.Fatalf("Failed to remove the original file: %v", err)
	}

	fmt.Printf("File successfully processed. Output at: %s\n", destinationFile)
}
