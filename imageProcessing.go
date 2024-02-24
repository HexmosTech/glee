package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/net/html"
)

func getImageFromPost(postJSON string) ([]string, error) {
	var post map[string]interface{}
	if err := json.Unmarshal([]byte(postJSON), &post); err != nil {
		log.Error("Error during JSON unmarshaling:", err)
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

	imageList := []string{}
	z := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "img" {
				for _, attr := range t.Attr {
					if attr.Key == "src" {
						imageList = append(imageList, attr.Val)
						break
					}
				}
			}
		}
	}

	return imageList, nil

}

func imageToHash(image string) (string, string, error) {
	var tp string // Temporary file path
	var err error
	var fileExtension string
	var hashValue string

	if strings.HasPrefix(image, "http://") || strings.HasPrefix(image, "https://") {

		iext := filepath.Ext(image)
		// fmt.Println("iext", iext)
		tempDir := getTempDir()

		tp = filepath.Join(tempDir, "img"+iext)
		resp, err := http.Get(image)
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
		fileExtension = iext
		hashValue, err = sha256Sum(tp)
	} else {
		// fmt.Println("Downloading image:", image)
		hashValue, err = sha256Sum(image)
		fileExtension = filepath.Ext(image)
	}

	if err != nil {
		return "", "", err
	}

	imageName := hashValue + fileExtension
	return imageName, tp, nil
}

func getTempDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("TEMP")
	}
	return "/tmp/"
}
