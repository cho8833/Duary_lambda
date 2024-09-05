package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"net/http"
	"sync"
)

var lock = &sync.Mutex{}

type CacheClient struct {
	client         *http.Client
	dynamoDBClient *dynamodb.Client
}

var httpClientInstance *CacheClient

func GetCacheClient() *CacheClient {
	if httpClientInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if httpClientInstance == nil {
			httpClientInstance = &CacheClient{}
		}
	}
	return httpClientInstance
}
func (cacheClient *CacheClient) GetHttpClient() (*http.Client, error) {
	if cacheClient.client == nil {
		cacheClient.client = &http.Client{}
	}
	return cacheClient.client, nil
}

func (cacheClient *CacheClient) GetDynamoDBClient() (*dynamodb.Client, error) {
	if cacheClient.dynamoDBClient == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			return nil, err
		}
		cacheClient.dynamoDBClient = dynamodb.NewFromConfig(cfg)
	}
	return cacheClient.dynamoDBClient, nil
}
