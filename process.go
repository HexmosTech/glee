package main

import (
	"bytes"
	"net/http"

	img "github.com/OhYee/goldmark-image"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkHTML "github.com/yuin/goldmark/renderer/html"
)

var theme string
var md = goldmark.New()
var buf bytes.Buffer

func postToGhost(metadata map[string]interface{}, content string) {
	injectMultiTitles(metadata)
	if val, ok := metadata["code_hilite_theme"]; ok {
		theme = val.(string)
		log.Debug("Theme from markdown:", theme)
	} else {
		codeTheme := config.GetDefault("blog-configuration.CODE_HILITE_THEME", "").(string)
		if codeTheme != "" {
			theme = codeTheme
		} else {
			theme = "monokai" // Default theme if none is specified
		}
		log.Debug("Default/Global Configured theme:", theme)
	}
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			img.NewImg("image", nil),
			highlighting.NewHighlighting(
				highlighting.WithStyle(theme),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkHTML.WithHardWraps(),
			goldmarkHTML.WithUnsafe(),
		),
	)

	metadata = addBlogConfigurations(metadata)
	metadata["html"] = toHTML(content)
	token, err := getJWToken()
	if err != nil {
		log.Error("Failed generate jwt token: %v", err)
	} else {
		log.Debug("Generated jwt token")
	}
	headers := http.Header{}
	headers.Set("Authorization", "Ghost "+token)
	post, err := getPostId(metadata["slug"].(string), headers)
	if err != nil {
		log.Error("Error: Unable to communicate with the Ghost Admin API. Please verify your Ghost configurations: %v\n", err)
		return
	}
	var pid, updated_at, htmlData, featureImage string
	if post == nil {
		log.Info("No post found for the given slug.")
	} else {
		pid, updated_at, htmlData, featureImage = post.ID, post.UpdatedAt, post.Mobiledoc, post.FeatureImage
	}
	if _, ok := metadata["feature_image"]; ok {
		uploadFeatureImage(metadata, token, featureImage)
	} else {
		log.Fatal("Feature image not provided")
		metadata["feature_image"] = ""
	}

	uploadedImages, err := uploadImages(token, htmlData)
	if err != nil {
		log.Error("Error uploading images:", err)
	} else {
		log.Info("Uploaded Blog Images")
	}
	result, err := replaceImageLinks(metadata, uploadedImages)
	if err != nil {
		log.Error("Error replacing image links:", err)
	}
	metadata["html"] = result

	postObj := metadata
	body := map[string]interface{}{
		"posts": []map[string]interface{}{postObj},
	}
	makeRequest(headers, body, pid, updated_at)

}
