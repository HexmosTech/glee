package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

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

func sha256Sum(filename string) (string, error) {
	h := sha256.New()

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	hashInBytes := h.Sum(nil)
	hash := hex.EncodeToString(hashInBytes)

	return hash, nil
}

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to communicate with the Ghost Admin API: %s", body)
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if len(data.Posts) > 0 {
		return &data.Posts[0], nil
	}

	return nil, errors.New("no posts found for the given slug")
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
		log.Error("Error marshaling body to JSON:", err)
		return
	}

	req, err := http.NewRequest(method, apiEndpoint, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Error("Error creating request:", err)
		return
	}
	req.Header = headers
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error executing request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errorMessage string
		errorBytes, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(errorBytes, &errorMessage); err != nil {
			errorMessage = string(errorBytes)
		}
		log.Error("Request failed with status: %s. Error message: %s\n", resp.Status, errorMessage)
		return
	}

	// Unmarshal the response
	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		log.Error("Error unmarshaling response:", err)
		return
	}

	// Log the result
	if pid == "" {
		log.Info("Created new post")
	} else {
		log.Info("Updated existing post based on slug")
	}
	fmt.Printf("Blog preview link: %s\n", responseData["posts"].([]interface{})[0].(map[string]interface{})["url"])
}
