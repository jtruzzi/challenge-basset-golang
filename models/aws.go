package models

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

func InitAwsServices() {
	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})
	db = dynamodb.New(awsSession)
}