package main

import (
	"regexp"
	"strings"
)

func uploadVideos(token, htmlData string) (map[string]string, error) {
	log.Info("Uploading Blog Videos ...")
	uploadedVideos := make(map[string]string)
	videoBackend := config.GetDefault("image-configuration.VIDEO_BACKEND", "").(string)
	log.Info("Videos Upload Backend...",videoBackend)
	blogVideoList := []string{}

	// Detect video tags and direct video links
	pattern := `<video[^>]*>.*?<source[^>]*src="([^"]+)"[^>]*>.*?</video>`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(buf.String(), -1)

	log.Info("Videos Upload Regex matches...",matches)

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
		hashValue, videoData, err := videoToHash(video) // Implement videoToHash function
		log.Info("Videos Value...",video)
		log.Info("Videos Upload Hash Value...",hashValue)
		log.Info("Videos Upload videoData...",videoData)
		if err != nil {
			log.Error("Error calculating hash:", err)
			return nil, err
		}

		var videoLink string
		if strings.HasPrefix(video, "http://") || strings.HasPrefix(video, "https://") {
			if videoBackend == "s3" {
				log.Info("Uploading video o S3...",videoLink,videoData,hashValue)
				videoLink, err = uploadToS3(videoData, hashValue)
					

			} else if videoBackend == "ghost" {
				videoLink, err = uploadToGhost(token, videoData, hashValue, blogVideoList)
			}
		} else {
			if videoBackend == "s3" {
				videoLink, err = uploadToS3(video, hashValue)
				log.Info("Uploading video o S3...",videoLink,videoData,hashValue)
			} else if videoBackend == "ghost" {
				videoLink, err = uploadToGhost(token, video, hashValue, blogVideoList)
			}
		}

		if err != nil {
			return nil, err
		}

		uploadedVideos[video] = videoLink
		
	}
	return uploadedVideos, nil
}
