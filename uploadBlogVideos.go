package main

import (
	"regexp"
)

func uploadVideos(token, htmlData string) (map[string]string, error) {
	uploadedVideos := make(map[string]string)
	videoBackend := config.GetDefault("video-configuration.VIDEO_BACKEND", "").(string)
	blogVideoList := []string{}

	// Detect video tags and direct video links
	pattern := `<video[^>]*>.*?<source[^>]*src="([^"]+)"[^>]*>.*?</video>`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(buf.String(), -1)

	// Get existing videos if backend is Ghost
	if videoBackend == "ghost" {
		var err error
		blogVideoList, err = getVideoFromPost(htmlData)
		if err != nil {
			log.Error("Error getting videos from post:", err)
		}
	}

	var mdlibVideos []string
	for _, m := range matches {
		if len(m) > 1 {
			mdlibVideos = append(mdlibVideos, m[1])
		}
	}
	// Upload videos and populate uploadedVideos map
	for _, video := range mdlibVideos {
		hashValue, _, err := videoToHash(video) // Implement videoToHash function
		if err != nil {
			log.Error("Error calculating hash:", err)
			return nil, err
		}

		var videoLink string

		if videoBackend == "s3" {
			videoLink, err = uploadToS3(video, hashValue)
		} else if videoBackend == "ghost" {
			videoLink, err = uploadToGhost(token, video, hashValue, blogVideoList)
		}

		if err != nil {
			return nil, err
		}

		uploadedVideos[video] = videoLink

	}
	return uploadedVideos, nil
}
