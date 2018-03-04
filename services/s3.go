package services

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"bytes"
	"fmt"
	"log"
	"net/url"
	"io/ioutil"
)

func SaveAttachmentToS3(fileName string, folder string, content []byte) string {
	// TODO: Upload to S3

	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})


	uploader := s3manager.NewUploader(awsSession)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("basset-mailing-gateway"),
		Key:    aws.String(folder + "/" + fileName),
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		fmt.Println("Error uploading file", err)
		return ""
	}

	log.Println("Successfully uploaded %s", fileName )
	return result.Location
}

func GetAttachmentFromS3(location string) ([]byte, error) {
	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})
	s3Svc := s3.New(awsSession)

	u,_ := url.Parse(location)
	fmt.Printf("proto: %q, bucket: %q, key: %q", u.Scheme, u.Host, u.Path)

	result, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("basset-mailing-gateway"),
		Key:    aws.String(u.Path),
	})

	if err != nil {
		log.Println(err)
	}

	defer result.Body.Close()
	return ioutil.ReadAll(result.Body)
}