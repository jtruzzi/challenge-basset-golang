package models

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"fmt"
)

type TicketRelease struct {
	ItemId string
	Released bool
	S3Url string
}

func GetTicketRelease(itemId string) (TicketRelease, error) {
	result, err := db.GetItem(&dynamodb.GetItemInput{
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

	item := TicketRelease{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		log.Println("Failed to unmarshal Record, %v", err)
		return TicketRelease{}, err
	}

	return item, nil
}


func CreateTicketRelease(itemId string, released bool, s3Url string) (TicketRelease, error) {
	item := TicketRelease{
		ItemId: itemId,
		Released: released,
		S3Url: s3Url,
	}

	av, err := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("TicketRelease"),
	}

	_, err = db.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return TicketRelease{}, err
	}

	return item, nil
}