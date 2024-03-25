package config

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"sync"
)

type DynamoDBOption struct {
	Client *dynamodb.Client
}

var lock = &sync.Mutex{}

var contextSingleton *DynamoDBOption

func GetDynamoDBOption() *DynamoDBOption {
	if contextSingleton == nil {
		lock.Lock()
		defer lock.Unlock()
	}
	if contextSingleton == nil {
		client, err := loadDynamoDBClient()
		if err != nil {
			fmt.Printf("failed to create DynamoDBOption : %s\n", err)
			return nil
		}
		contextSingleton = &DynamoDBOption{Client: client}
	}
	return contextSingleton
}

func loadDynamoDBClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}
