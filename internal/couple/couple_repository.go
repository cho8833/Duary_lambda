package couple

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	uuid2 "github.com/google/uuid"
)

const tableName = "Couple"

type Repository interface {
	SaveCouple(couple *Couple) (*Couple, error)
	GetSaveCoupleTransaction(couple *Couple) (*types.TransactWriteItem, error)
}

type RepositoryDynamoDB struct {
	client *dynamodb.Client
}

func NewCoupleRepository(client *dynamodb.Client) *RepositoryDynamoDB {
	return &RepositoryDynamoDB{client: client}
}

func (repository *RepositoryDynamoDB) SaveCouple(couple *Couple) (*Couple, error) {
	couple.Id = generateUID()
	item, err := attributevalue.MarshalMap(couple)
	if err != nil {
		return nil, err
	}

	_, err = repository.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Couple"),
		Item:      item,
	})

	if err != nil {
		return nil, err
	}

	return couple, nil
}

func (repository *RepositoryDynamoDB) GetSaveCoupleTransaction(couple *Couple) (*types.TransactWriteItem, error) {
	couple.Id = generateUID()
	item, err := attributevalue.MarshalMap(couple)
	if err != nil {
		return nil, err
	}
	transaction := &types.TransactWriteItem{Put: &types.Put{
		TableName: aws.String(tableName),
		Item:      item,
	}}

	return transaction, nil
}

func generateUID() *string {
	uuid := uuid2.New().String()
	return &uuid
}
