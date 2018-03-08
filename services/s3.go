package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"../models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func SaveAttachmentToS3(attachment models.Attachment) string {
	// TODO: Upload to S3

	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	uploader := s3manager.NewUploader(awsSession)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("basset-mailing-gateway"),
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
		Region: aws.String("us-east-1"),
	})
	s3Svc := s3.New(awsSession)

	u, _ := url.Parse(location)
	fmt.Printf("proto: %q, bucket: %q, key: %q", u.Scheme, u.Host, u.Path)

	result, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("basset-mailing-gateway"),
		Key:    aws.String(u.Path),
	})

	if err != nil {
		log.Println(err)
	}

	defer result.Body.Close()
	content, _ := ioutil.ReadAll(result.Body)

	return models.Attachment{
		Mime:    "application/pdf",
		Path:    u.Path,
		Content: content,
	}, nil
}
