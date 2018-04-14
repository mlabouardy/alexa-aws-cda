package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"github.com/rs/xid"
)

type Service struct {
	Category  string
	Questions []Question
}

type Question struct {
	Question string
	Answers  map[string]string
	Correct  string
}

type DBItem struct {
	ID       string
	Category string
	Question string
	Answers  map[string]string
	Correct  string
}

func insertToDynamoDB(cfg aws.Config, category string, question Question) error {
	tableName := os.Getenv("TABLE_NAME")

	item := DBItem{
		ID:       xid.New().String(),
		Category: category,
		Question: question.Question,
		Answers:  question.Answers,
		Correct:  question.Correct,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println(err)
		return err
	}

	svc := dynamodb.New(cfg)
	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      av,
	})
	_, err = req.Send()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadFile("questions.json")
	if err != nil {
		log.Fatal(err)
	}

	services := make([]Service, 0)
	json.Unmarshal(raw, &services)

	for _, service := range services {
		fmt.Printf("AWS Service: %s\n", service.Category)
		fmt.Printf("\tQuestions: %d\n", len(service.Questions))
		for _, question := range service.Questions {
			err = insertToDynamoDB(cfg, service.Category, question)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
