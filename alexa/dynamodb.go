package main

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

type DBItem struct {
	ID       string
	Category string
	Question string
	Answers  map[string]string
	Correct  string
}

func getQuestions(category string) ([]DBItem, error) {
	items := make([]DBItem, 0)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return items, err
	}

	svc := dynamodb.New(cfg)
	req := svc.ScanRequest(&dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("TABLE_NAME")),
		FilterExpression: aws.String("Category = :c"),
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":c": dynamodb.AttributeValue{
				S: aws.String(category),
			},
		},
	})
	result, err := req.Send()
	if err != nil {
		return items, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return items, err
	}

	return items, nil
}
