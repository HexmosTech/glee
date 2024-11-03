package main

import (
	"bytes"
	"errors"
	"strings"

	"golang.org/x/net/html"
)

func toHTML(markdown string) string {
	start := "<!--kg-card-begin: html-->"
	end := "<!--kg-card-end: html-->"
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		log.Error("Error converting markdown to html: %v", err)
	}

	return start + buf.String() + end
}

func replaceMediaLinks(metadata map[string]interface{}, imgMap, videoMap map[string]string) (string, error) {
	htmlStr, ok := metadata["html"].(string)
	if !ok {
		return "", errors.New("metadata does not contain a string under the 'html' key")
	}

	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return "", err
	}

	// Function to replace both image and video links
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i, attr := range n.Attr {
				switch {
				case n.Data == "img" && attr.Key == "src":
					fallthrough
				case n.Data == "a" && attr.Key == "href":
					if newSrc, ok := imgMap[attr.Val]; ok {
						n.Attr[i].Val = newSrc
					}
				case n.Data == "source" && attr.Key == "src":
					if newSrc, ok := videoMap[attr.Val]; ok {
						n.Attr[i].Val = newSrc
					}
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(doc)

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
