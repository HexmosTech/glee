package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/pelletier/go-toml"

	// "github.com/russross/blackfriday/v2"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var config *toml.Tree

type Post struct {
	ID           string `json:"id"`
	UpdatedAt    string `json:"updated_at"`
	Mobiledoc    string `json:"mobiledoc"`
	FeatureImage string `json:"feature_image"`
}

type ResponseData struct {
	Posts []Post `json:"posts"`
}

var postsApiBase string

func getPostId(slug string, headers http.Header) (*Post, error) {
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)
	ghostUrl := config.GetDefault("ghost-configuration.GHOST_URL", "").(string)

	if ghostVersion == "v5" {
		postsApiBase = ghostUrl + "/api/admin/posts/"
	} else {
		postsApiBase = ghostUrl + "/api/" + ghostVersion + "/admin/posts/"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", postsApiBase+"slug/"+slug+"/", nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("Unable to communicate with the Ghost Admin API: %s", body)
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if len(data.Posts) > 0 {
		return &data.Posts[0], nil
	}

	return nil, errors.New("No posts found for the given slug")
}

func toHTML(markdown string) string {
	start := "<!--kg-card-begin: html-->"
	end := "<!--kg-card-end: html-->"

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // Includes table, fenced code, and code highlight extensions
			extension.Table,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // Enable automatic heading IDs for TOC
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(), // Enable hard wrapping for better formatting
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		panic(err)
	}
	// fmt.Printf("HTML: %s\n", buf.String())
	return start + buf.String() + end // Return the generated HTML string
}

func initLogging(debug bool) {
	log.SetOutput(os.Stdout)

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// You can customize the log formatter if needed
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func addBlogConfigurations(meta map[string]interface{}) map[string]interface{} {
	// Check if config is nil
	if config == nil {
		log.Fatal("Configuration is not initialized. Call viewTOMLFile to initialize it.")
	}

	globalSidebarTOC := config.GetDefault("blog-configuration.SIDEBAR_TOC", "").(bool)
	globalFeatured := config.GetDefault("blog-configuration.FEATURED", "").(bool)
	globalStatus := config.GetDefault("blog-configuration.STATUS", "").(string)

	// Assuming meta is a map[string]interface{}, add or update relevant fields
	fmt.Printf("SIDEBAR_TOC: %v\n", globalSidebarTOC)
	fmt.Printf("FEATURED: %v\n", globalFeatured)
	fmt.Printf("STATUS: %v\n", globalStatus)

	return meta
}

func viewTOMLFile() {
	configPath := filepath.Join(os.Getenv("HOME"), ".glee.toml")
	var err error
	config, err = toml.LoadFile(configPath)
	if err != nil {
		getTOMLFile(configPath)
		os.Exit(0)
	}
}

func getTOMLFile(configPath string) {
	fmt.Printf("The configuration file at %s was not found.\n", configPath)

	var configResponse string
	fmt.Print("Would you like me to create the configuration file? (yes/no): ")
	fmt.Scanln(&configResponse)

	if configResponse == "yes" || configResponse == "y" {
		// Your existing code to create the configuration file
		// ...

	} else {
		os.Exit(0)
	}
}

func makeRequest(headers http.Header, body map[string]interface{}, pid string, updated_at string) {
	var method, apiEndpoint string
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)
	ghostUrl := config.GetDefault("ghost-configuration.GHOST_URL", "").(string)

	if ghostVersion == "v5" {
		postsApiBase = ghostUrl + "/api/admin/posts/"
	} else {
		postsApiBase = ghostUrl + "/api/" + ghostVersion + "/admin/posts/"
	}
	if pid == "" {
		// fmt.Printf("%#v\n", body)
		method = http.MethodPost
		apiEndpoint = postsApiBase + "?source=html"
	} else {
		method = http.MethodPut
		body["posts"].([]map[string]interface{})[0]["updated_at"] = updated_at
		apiEndpoint = postsApiBase + pid + "?source=html"
	}
	

	client := &http.Client{}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling body to JSON:", err)
		return
	}

	req, err := http.NewRequest(method, apiEndpoint, bytes.NewBuffer(bodyJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header = headers
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	defer resp.Body.Close()

if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
    var errorMessage string
    errorBytes, _ := ioutil.ReadAll(resp.Body) // Read the response body
    if err := json.Unmarshal(errorBytes, &errorMessage); err != nil {
        // If the error message is not JSON, just use the raw response body
        errorMessage = string(errorBytes)
    }
    fmt.Printf("Request failed with status: %s. Error message: %s\n", resp.Status, errorMessage)
    return
}


	// Unmarshal the response
	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	// Log the result
	if pid == "" {
		fmt.Println("Created new post")
	} else {
		fmt.Println("Updated existing post based on slug")
	}
	fmt.Printf("Blog preview link: %s\n", responseData["posts"].([]interface{})[0].(map[string]interface{})["url"])
}

func postToGhost(metadata map[string]interface{}, content string) {
	// Add configurations to metadata
	addBlogConfigurations(metadata)
	metadata["html"] = toHTML(content)
	token, err := getJWToken()
	if err != nil {
		log.Fatalf("Failed generate jwt token: %v", err)
	}
	headers := http.Header{}
	headers.Set("Authorization", "Ghost "+token)
	post, err := getPostId(metadata["slug"].(string), headers)
	if err != nil {
		fmt.Printf("Error: Unable to communicate with the Ghost Admin API. Please verify your Ghost configurations: %v\n", err)
		return
	}
	
	// fmt.Printf("Post ID: %s\nUpdated At: %s\nHTML Data: %s\nFeature Image: %s\n",pid, updated_at, html_data, feature_image)
	var pid, updated_at string
	if post == nil {
		fmt.Println("No post found for the given slug.")
	} else {
		pid, updated_at = post.ID, post.UpdatedAt
	}
	postObj := metadata
	body := map[string]interface{}{
		"posts": []map[string]interface{}{postObj},
	}
	makeRequest(headers, body, pid, updated_at)

	// Print YAML front matter
	// fmt.Printf("HTML Content: %s\n", metadata["html"])
	// yamlData, err := yaml.Marshal(metadata)
	// if err != nil {
	// 	log.Fatalf("Failed to convert metadata to YAML: %v", err)
	// }
	// fmt.Println(string(yamlData))

	// Print Markdown content
	// fmt.Println("Markdown Content:")
	// fmt.Println(content)

}

func getJWToken() (string, error) {

	ghostAdminKey := config.GetDefault("ghost-configuration.ADMIN_API_KEY", "").(string)
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)

	parts := strings.Split(ghostAdminKey, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid admin key format")
	}
	id, secret := parts[0], parts[1]

	var audValue string
	if ghostVersion == "v5" {
		audValue = "/admin/"
	} else {
		audValue = "/" + ghostVersion + "/admin/"
	}

	iat := time.Now().Unix()
	exp := iat + 5*60 // expires in  5 minutes

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": iat,
		"exp": exp,
		"aud": audValue,
	})

	token.Header["kid"] = id

	secretBytes, err := hex.DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to decode secret: %v", err)
	}

	signedToken, err := token.SignedString(secretBytes)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}

// func newMarkdown() goldmark.Markdown {
// 	return goldmark.New(
// 		goldmark.WithExtensions(
// 			extension.GFM, // Includes table, fenced code, and code highlight extensions
// 			extension.Table,
// 		),
// 		goldmark.WithParserOptions(
// 			parser.WithAutoHeadingID(), // Enable automatic heading IDs for TOC
// 		),
// 		goldmark.WithRendererOptions(
// 			html.WithHardWraps(), // Enable hard wrapping for better formatting
// 		),
// 	)
// }

func main() {
	// newMarkdown()
	initLogging(true)
	viewTOMLFile()
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run glee.go <markdown_file>")
		os.Exit(1)
	}

	// Read the markdown file from the command line argument
	filePath := os.Args[1]
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
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

	// Parse the YAML front matter into a Go struct or map
	var metadata map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlFrontMatter), &metadata)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// Call the postToGhost function with metadata and content
	postToGhost(metadata, markdownContent)
}
