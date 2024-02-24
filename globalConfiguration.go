package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

func loadGlobalConfig() {
	configPath := filepath.Join(os.Getenv("HOME"), ".glee.toml")
	var err error
	config, err = toml.LoadFile(configPath)
	if err != nil {
		getTomlFile(configPath)
		os.Exit(0)
	}
}

func checkConfigurationsExist(log *logrus.Logger) {
	checkField := func(section, field, expectedType string) {
		value, ok := config.GetPath([]string{section, field}).(string)
		if !ok || value == "" {
			log.Errorf("Error: Include %s %s in the configuration", section, field)
			os.Exit(1)

		}
	}

	checkField("ghost-configuration", "ADMIN_API_KEY", "string")
	checkField("ghost-configuration", "GHOST_URL", "string")
	checkField("ghost-configuration", "GHOST_VERSION", "string")

	checkField("image-configuration", "IMAGE_BACKEND", "string")

	if imageBackend := config.GetPath([]string{"image-configuration", "IMAGE_BACKEND"}).(string); imageBackend == "s3" {
		checkField("aws-s3-configuration", "ACCESS_KEY_ID", "string")
		checkField("aws-s3-configuration", "SECRET_ACCESS_KEY", "string")
		checkField("aws-s3-configuration", "BUCKET_NAME", "string")
		checkField("aws-s3-configuration", "S3_BASE_URL", "string")
	}
}

func printConfiguration() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	configPath := filepath.Join(usr.HomeDir, ".glee.toml")
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	fmt.Printf("Configuration path: %s\n", configPath)
	fmt.Println("-------------------")
	fmt.Println(string(content))
}

func getTomlFile(configPath string) {
	log.Warning("The configuration file not found at.", configPath)
	reader := bufio.NewReader(os.Stdin)
	log.Info("Would you like me to create the configuration file? (yes/no): ")
	configResponse, _ := reader.ReadString('\n')
	configResponse = configResponse[:len(configResponse)-1] // Remove newline

	if configResponse == "yes" || configResponse == "y" {
		url := "https://raw.githubusercontent.com/HexmosTech/glee/main/.glee.toml"
		response, err := http.Get(url)
		if err != nil {
			log.Error("Failed to create the configuration file: %v", err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			content, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Error("Failed to read response content: %v", err)
				return
			}
			err = ioutil.WriteFile(configPath, content, 0644)
			if err != nil {
				log.Error("Failed to create the configuration file: %v", err)
				return
			}
			log.Info("Created the configuration file in", configPath)
			msg := fmt.Sprintf("Include the Ghost configurations in the file located at %s", configPath)
			log.Info(msg)
		} else {
			log.Error("Failed to create the configuration file")
		}
	}
}
