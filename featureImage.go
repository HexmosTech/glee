package main

import "path/filepath"

func uploadFeatureImage(meta map[string]interface{}, token string, featureImage string) {
	if featureImagePath, ok := meta["feature_image"].(string); ok {
		hashValue, err := sha256Sum(featureImagePath)
		if err != nil {
			log.Error("Error calculating SHA-256 sum:", err)
			return
		}

		fileExtension := filepath.Ext(featureImagePath)
		imageName := hashValue + fileExtension
		imageBackend := config.GetDefault("image-configuration.IMAGE_BACKEND", "").(string)

		if imageBackend == "s3" {
			log.Debug("Uploading feature image to AWS S3")
			meta["feature_image"], err = uploadToS3(featureImagePath, imageName)
			if err != nil {
				log.Error("Error uploading file to S3:", err)
			} else {
				log.Info("Images uploading to S3")
			}
		} else {
			log.Debug("Uploading feature image to Ghost Database")
			var featureImgList []string
			if featureImage != "" {
				featureImgList = []string{featureImage}
			} else {
				featureImgList = []string{}
			}
			meta["feature_image"], err = uploadToGhost(token, featureImagePath, imageName, featureImgList)
			if err != nil {
				log.Error("Error uploading feature image to Ghost:", err)
			}

		}

		log.Info("Uploaded feature image")
	} else {
		log.Info("Feature image not provided")
	}
}
