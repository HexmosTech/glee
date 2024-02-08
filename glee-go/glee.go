package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ghodss/yaml"
)

func main() {
	// Read the markdown file
	content, err := ioutil.ReadFile("sample_post.md")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Convert to string
	contentStr := string(content)

	// Find the index of the second delimiter to isolate the YAML front matter
	index := strings.Index(contentStr, "\n\n---\n\n")
	if index == -1 {
		log.Fatal("Could not find YAML front matter delimiter")
	}

	// Extract the YAML front matter and the rest of the content
	yamlFrontMatter := contentStr[:index]
	markdownContent := contentStr[index+len("\n\n---\n\n"):]

	// Print the YAML front matter and the Markdown content
	fmt.Println("YAML Front Matter:")
	fmt.Println(yamlFrontMatter)
	fmt.Println("Markdown Content:")
	fmt.Println(markdownContent)

	// Parse the YAML front matter into a Go struct or map
	var metadata map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlFrontMatter), &metadata)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// Now you can work with the metadata map
	fmt.Printf("Parsed Metadata: %+v\n", metadata)
}
