package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoDB *dynamodb.DynamoDB

func InitDB() {
	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	DynamoDB = dynamodb.New(awsSession)
}
