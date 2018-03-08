package models

import (
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"../db"
)

type TicketRelease struct {
	ItemId   string
	Released bool
	S3Url    string
}

func GetTicketRelease(itemId string) (TicketRelease, error) {
	result, err := db.DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("TicketRelease"),
		Key: map[string]*dynamodb.AttributeValue{
			"ItemId": {
				S: aws.String(itemId),
			},
		},
	})
	if err != nil {
		log.Println(err.Error())
		return TicketRelease{}, err
	}

	item := TicketRelease{ItemId: itemId}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		log.Println("Failed to unmarshal Record", err)
		return item, err
	}

	return item, nil
}

func (item *TicketRelease) Save() error {
	av, err := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("TicketRelease"),
	}

	_, err = db.DynamoDB.PutItem(input)

	if err != nil {
		log.Println("Got error calling PutItem:")
		log.Println(err.Error())
		return err
	}

	return nil
}
