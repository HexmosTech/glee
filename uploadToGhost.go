package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

func uploadToGhost(token, imageData, hashName string, blogImageList []string) (string, error) {
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)
	ghostUrl := config.GetDefault("ghost-configuration.GHOST_URL", "").(string)
	var postsApiBase string
	if ghostVersion == "v5" {
		postsApiBase = ghostUrl + "/api/admin/images/upload/"
	} else {
		postsApiBase = ghostUrl + "/api/" + ghostVersion + "/admin/images/upload/"
	}

	for _, name := range blogImageList {
		hashValue := strings.Split(hashName, ".")[0]
		filename := filepath.Base(name)
		filenameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
		if hashValue == filenameWithoutExt {
			log.Debug("The image already exists and is being reused: ", name)
			return name, nil
		}
	}

	// Open the file for reading
	file, err := os.Open(imageData)

	if err != nil {
		return "", errors.New("Failed to open image file: " + err.Error())
	}
	defer file.Close()
	// Create a buffer to hold the multipart encoded form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	contentType := mime.TypeByExtension(filepath.Ext(imageData))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	part, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{`form-data; name="file"; filename="` + hashName + `"`},
		"Content-Type":        []string{contentType},
	})

	if err != nil {
		return "", errors.New("Failed to create form file part: " + err.Error())
	}
	_, err = io.Copy(part, file)

	if err != nil {
		return "", errors.New("Failed to copy image data to form part: " + err.Error())
	}

	err = writer.WriteField("ref", hashName)
	if err != nil {
		return "", errors.New("Failed to write ref field: " + err.Error())
	}

	err = writer.Close()

	if err != nil {
		return "", errors.New("Failed to close multipart writer: " + err.Error())
	}

	req, err := http.NewRequest("POST", postsApiBase, body)

	if err != nil {
		return "", errors.New("Failed to create HTTP request: " + err.Error())
	}
	req.Header.Set("Authorization", "Ghost "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.New("Failed to send HTTP request: " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return "", errors.New("Failed to upload image: HTTP status " + resp.Status)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", errors.New("Failed to decode JSON response: " + err.Error())
	}

	imageLink := result["images"].([]interface{})[0].(map[string]interface{})["url"].(string)

	return imageLink, nil
}
