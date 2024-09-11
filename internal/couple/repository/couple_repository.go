package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/cho8833/duary_lambda/internal/couple/model"
	uuid2 "github.com/google/uuid"
)

type CoupleRepository interface {
	SaveCouple(couple *model.Couple) (*model.Couple, error)
}

type CoupleRepositoryDynamoDB struct {
	client *dynamodb.Client
}

func NewCoupleRepository(client *dynamodb.Client) *CoupleRepositoryDynamoDB {
	return &CoupleRepositoryDynamoDB{client: client}
}

func (repository *CoupleRepositoryDynamoDB) SaveCouple(couple *model.Couple) (*model.Couple, error) {
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
