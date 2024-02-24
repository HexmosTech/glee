package main

import (
	"bytes"
	"mime"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadToS3(localFilePath, s3FilePath string) (string, error) {
	accessKeyID := config.GetDefault("aws-s3-configuration.ACCESS_KEY_ID", "").(string)
	secretAccessKey := config.GetDefault("aws-s3-configuration.SECRET_ACCESS_KEY", "").(string)
	bucketName := config.GetDefault("aws-s3-configuration.BUCKET_NAME", "").(string)
	s3BaseUrl := config.GetDefault("aws-s3-configuration.S3_BASE_URL", "").(string)
	region := config.GetDefault("aws-s3-configuration.REGION", "").(string)

	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region), Credentials: creds},
	)
	if err != nil {
		log.Error("Error creating session:", err)
		return "", err
	}

	svc := s3.New(sess)

	mimeType, err := getMimeType(localFilePath)
	if err != nil {
		log.Error("Error determining MIME type:", err)
		return "", err
	}

	file, err := os.Open(localFilePath)
	if err != nil {
		log.Error("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Error("Error getting file stats:", err)
		return "", err
	}
	size := stat.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(s3FilePath),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(mimeType),
	}

	_, err = svc.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fallthrough
			default:
				log.Error(aerr.Error())
			}
		} else {
			log.Error(err.Error())
		}
		return "", err
	}

	s3URL := s3BaseUrl + s3FilePath
	log.Debug("Image s3 url", s3URL)
	return s3URL, nil
}

func getMimeType(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	return mime.TypeByExtension(ext), nil
}
