package couple

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	uuid2 "github.com/google/uuid"
)

type Repository interface {
	SaveCouple(couple *Couple) (*Couple, error)
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

func generateUID() string {
	uuid := uuid2.New()
	return uuid.String()
}
