package services

import (
	"bytes"
	"io/ioutil"
	"log"
	"../models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"path/filepath"
)

func SaveAttachmentToS3(attachment models.Attachment) string {
	// TODO: Upload to S3

	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_S3_REGION")),
	})

	uploader := s3manager.NewUploader(awsSession)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mailing-gateway"),
		Key:    aws.String(attachment.Path),
		Body:   bytes.NewReader(attachment.Content),
	})
	if err != nil {
		log.Println("Error uploading file", err)
		return ""
	}

	log.Println("Successfully uploaded %v", attachment.Path)
	return result.Location
}

func GetAttachmentFromS3(location string) (models.Attachment, error) {
	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_S3_REGION")),
	})
	s3Svc := s3.New(awsSession)

	fileName := filepath.Base(location)

	result, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("mailing-gateway"),
		Key:    aws.String(fileName),
	})

	if err != nil {
		log.Println(err)
	}

	defer result.Body.Close()
	content, _ := ioutil.ReadAll(result.Body)

	return models.Attachment{
		Mime:    "application/pdf",
		Path:    fileName,
		Content: content,
	}, nil
}
