package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func getVideoFromPost(postJSON string) ([]string, error) {
	var post map[string]interface{}
	if err := json.Unmarshal([]byte(postJSON), &post); err != nil {
		log.Error("Error during JSON unmarshaling:", err)
		return nil, err
	}
	htmlContent := ""
	if cards, ok := post["cards"].([]interface{}); ok {
		if len(cards) > 0 {
			if cardData, ok := cards[0].([]interface{}); ok && len(cardData) > 1 {
				if cardHtml, ok := cardData[1].(map[string]interface{}); ok {
					if htmlStr, ok := cardHtml["html"].(string); ok {
						htmlContent = htmlStr
					}
				}
			}
		}
	}

	videoList := []string{}
	z := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "video" {
				for _, attr := range t.Attr {
					if attr.Key == "src" {
						videoList = append(videoList, attr.Val)
						break
					}
				}
			}
		}
	}
	return videoList, nil
}

// videoToHash downloads a video from a URL, hashes its content, and saves it locally.
func videoToHash(video string) (string, string, error) {

	var tp string // Temporary file path
	var err error
	var fileExtension string
	var hashValue string
	if strings.HasPrefix(video, "http://") || strings.HasPrefix(video, "https://") {
		vext := filepath.Ext(video)
		tempDir := getTempDir()

		tp = filepath.Join(tempDir, "vid"+vext)
		resp, err := http.Get(video)
		if err != nil {
			return "", "", err
		}
		defer resp.Body.Close()

		outFile, err := os.Create(tp)
		if err != nil {
			return "", "", err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			return "", "", err
		}

		// Create the "videos" directory if it doesn't exist
		if err = os.MkdirAll("videos", os.ModePerm); err != nil {
			return "", "", err
		}

		// Save the video with a unique name based on its hash
		hashValue, err = sha256Sum(tp)
		if err != nil {
			return "", "", err
		}
		fileExtension = vext

	} else {
		hashValue, err = sha256Sum(video)
		fileExtension = filepath.Ext(video)
	}
	if err != nil {
		return "", "", err
	}

	videoName := hashValue + fileExtension
	return videoName, tp, nil
}
