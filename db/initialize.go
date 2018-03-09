package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

var DynamoDB *dynamodb.DynamoDB

func InitDB() {
	var awsSession, _ = session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DYNAMODB_REGION")),
	})
	DynamoDB = dynamodb.New(awsSession)
}
