package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"sync"
)

type DynamoDBOption struct {
	Client *dynamodb.Client
}

var lock = &sync.Mutex{}

var contextSingleton *DynamoDBOption

func GetDynamoDBOption() (*DynamoDBOption, error) {
	if contextSingleton == nil {
		lock.Lock()
		defer lock.Unlock()
	}
	if contextSingleton == nil {
		client, err := loadDynamoDBClient()
		if err != nil {
			log.Printf("failed to create DynamoDBOption : %s\n", err)
			return nil, err
		}
		contextSingleton = &DynamoDBOption{Client: client}
	}
	return contextSingleton, nil
}

func loadDynamoDBClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}
