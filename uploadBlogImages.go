package main

import (
	"regexp"
	"strings"
)

func uploadImages(token, htmlData string) (map[string]string, error) {
	uploadedImages := make(map[string]string)
	imageBackend := config.GetDefault("image-configuration.IMAGE_BACKEND", "").(string)
	blogImageList := []string{}

	pattern := `<img[^>]+src="([^"]+)"[^>]*>`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(buf.String(), -1)
	if imageBackend == "ghost" {
		var err error
		blogImageList, err = getImageFromPost(htmlData)
		if err != nil {
			log.Error("Error getting images from post:", err)

		}
	}
	var mdlibImages []string
	for _, m := range matches {
		if len(m) > 1 {
			mdlibImages = append(mdlibImages, m[1])
		}
	}

	for _, image := range mdlibImages {
		hashValue, imageData, err := imageToHash(image) // Implement imageToHash function
		if err != nil {
			log.Error("Error calculating hash:", err)
			return nil, err
		}

		var imageLink string
		if strings.HasPrefix(image, "http://") || strings.HasPrefix(image, "https://") {
			if imageBackend == "s3" {
				imageLink, err = uploadToS3(imageData, hashValue)
			} else if imageBackend == "ghost" {
				imageLink, err = uploadToGhost(token, imageData, hashValue, blogImageList)
			}
		} else {
			if imageBackend == "s3" {
				log.Debug("Uploading image to AWS S3")
				imageLink, err = uploadToS3(image, hashValue)
			} else if imageBackend == "ghost" {
				imageLink, err = uploadToGhost(token, image, hashValue, blogImageList)
				log.Debug("Uploading image to Ghost Database")
			}
		}

		if err != nil {
			return nil, err
		}

		uploadedImages[image] = imageLink

	}
	return uploadedImages, nil
}
