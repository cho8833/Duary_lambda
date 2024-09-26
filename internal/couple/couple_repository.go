package couple

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	uuid2 "github.com/google/uuid"
	"log"
)

const tableName = "Couple"

type Repository interface {
	SaveCouple(couple *Couple) (*Couple, error)
	GetSaveCoupleTransaction(couple *Couple) (*types.TransactWriteItem, error)
	FindById(id *string) (*Couple, error)
	FindByCoupleCode(coupleCode *string) ([]Couple, error)
}

type RepositoryDynamoDB struct {
	client *dynamodb.Client
}

func NewCoupleRepository(client *dynamodb.Client) *RepositoryDynamoDB {
	return &RepositoryDynamoDB{client: client}
}

func (repository *RepositoryDynamoDB) SaveCouple(couple *Couple) (*Couple, error) {
	if couple.Id == nil {
		couple.Id = repository.generateUID()
	}
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

func (repository *RepositoryDynamoDB) UpdateCouple(req *UpdateCoupleReq) (*Couple, error) {
	updateExpr, err := repository.updateCoupleExpression(req)
	if err != nil {
		return nil, err
	}
	response, err := repository.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       repository.getKey(*req.Id),
		ExpressionAttributeValues: updateExpr.Values(),
		ExpressionAttributeNames:  updateExpr.Names(),
		UpdateExpression:          updateExpr.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		log.Printf("failed to update item. Req: %+v, error: %s", req, err.Error())
		return nil, err
	}
	result := &Couple{}
	_ = attributevalue.UnmarshalMap(response.Attributes, result)
	return result, nil
}

func (repository *RepositoryDynamoDB) FindByCoupleCode(coupleCode *string) ([]Couple, error) {
	filterEx := expression.Name("code").Equal(expression.Value(*coupleCode))

	expr, err := expression.NewBuilder().WithFilter(filterEx).Build()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	result, err := repository.client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	var couples []Couple
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &couples); err != nil {
		return nil, err
	}
	return couples, nil
}

func (repository *RepositoryDynamoDB) FindById(id *string) (*Couple, error) {
	response, err := repository.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       repository.getKey(*id),
	})
	if err != nil {
		return nil, err
	}
	result := &Couple{}
	_ = attributevalue.UnmarshalMap(response.Item, result)
	return result, nil
}

func (repository *RepositoryDynamoDB) GetSaveCoupleTransaction(couple *Couple) (*types.TransactWriteItem, error) {
	if couple.Id == nil {
		couple.Id = repository.generateUID()
	}
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

func (repository *RepositoryDynamoDB) generateUID() *string {
	uuid := uuid2.New().String()
	return &uuid
}

func (repository *RepositoryDynamoDB) getKey(id string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: id},
	}
}

func (repository *RepositoryDynamoDB) updateCoupleExpression(req *UpdateCoupleReq) (*expression.Expression, error) {
	update := expression.UpdateBuilder{}
	if req.IsConnected != nil {
		update = update.Set(expression.Name("isConnected"), expression.Value(*req.IsConnected))
	}

	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		log.Printf("failed to build update expression : %s", err.Error())
		return nil, err
	}
	return &expr, nil
}
