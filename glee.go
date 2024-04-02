package main

import (
	"fmt"

	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/jessevdk/go-flags"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

var config *toml.Tree
var log = logrus.New()
var filePath string

var opts struct {
	ShowConfig bool   `short:"c" long:"config" description:"Show glee configuration global file"`
	Debug      bool   `short:"d" long:"debug" description:"debug mode"`
	File       string `description:"Markdown file to process"`
	Help       bool   `short:"h" long:"help" description:"Show this help message"`
	Version    bool   `short:"v" long:"version" description:"Show version"`
	DownloadImage bool `short:"i" long:"download-image" description:"Download images from markdown file"`
}

var flagParser = flags.NewParser(&opts, flags.Default)

var version string

func initializeLogger() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}
func setupLogLevel() {
	if opts.Debug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

func parseArgs() {

	if len(version) == 0 {
		version = "vUnset"
	}
	if opts.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	if opts.Help {
		flagParser.WriteHelp(os.Stdout)
		os.Exit(0)
	}
}

func main() {
	initializeLogger()
	flagParser.Usage = "Usage: glee <markdown_file_path>"
	args, err := flagParser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	parseArgs()
	setupLogLevel()
	loadGlobalConfig()
	if opts.ShowConfig {
		printConfiguration()
		return
	}
	checkConfigurationsExist(log)
	if len(args) == 1 {
		filePath = args[0]
		content, err := os.ReadFile(filePath)

		if err != nil {
			log.Fatalf("Failed to read file %s: %v", filePath, err)
		}

		contentStr := string(content)

		index := strings.Index(contentStr, "\n---\n")
		if index == -1 {
			log.Fatal("Could not find YAML delimiter")
		}

		yamlFrontMatter := contentStr[:index]
		markdownContent := contentStr[index+len("\n---\n"):]

		var metadata map[string]interface{}
		err = yaml.Unmarshal([]byte(yamlFrontMatter), &metadata)
		if err != nil {
			log.Fatalf("Failed to parse YAML: %v", err)
		}

		postToGhost(metadata, markdownContent)
	} else {
		log.Fatal("Please provide a markdown file.")
	}
}
