package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type Client struct {
	ClientId       string
	Name           string
	MandrillApiKey string
}

func GetClient(clientId string) (Client, error) {
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Client"),
		Key: map[string]*dynamodb.AttributeValue{
			"ClientId": {
				S: aws.String(clientId),
			},
		},
	})
	if err != nil {
		log.Println(err.Error())
		return Client{}, err
	}

	item := Client{ClientId: clientId}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		log.Println("Failed to unmarshal Record", err)
		return item, err
	}

	return item, nil
}
