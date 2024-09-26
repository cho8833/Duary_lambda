package connectcouple

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

const tableName = "ConnectCoupleSession"

type SessionRepository interface {
	SaveSession(session *Session) (*Session, error)
	FindByCoupleCode(coupleCode *string) ([]Session, error)
	FindByConnectionId(connectionId *string) ([]Session, error)
	DeleteByConnectionId(key *string) error
}

type SessionRepositoryDynamoDB struct {
	client *dynamodb.Client
}

func NewConnectCoupleRepository(client *dynamodb.Client) *SessionRepositoryDynamoDB {
	return &SessionRepositoryDynamoDB{client: client}
}

func (repository *SessionRepositoryDynamoDB) SaveSession(session *Session) (*Session, error) {
	item, err := attributevalue.MarshalMap(session)
	if err != nil {
		return nil, err
	}
	_, err = repository.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (repository *SessionRepositoryDynamoDB) FindByCoupleCode(coupleCode *string) ([]Session, error) {
	filterEx := expression.Name("coupleCode").Equal(expression.Value(*coupleCode))

	expr, err := expression.NewBuilder().WithFilter(filterEx).Build()
	if err != nil {
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

	var sessions []Session
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

func (repository *SessionRepositoryDynamoDB) FindByConnectionId(connectionId *string) ([]Session, error) {
	keyEx := expression.Key("connectionId").Equal(expression.Value(types.AttributeValueMemberS{Value: *connectionId}))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}
	result, err := repository.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		return nil, err
	}

	var sessions []Session
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

func (repository *SessionRepositoryDynamoDB) DeleteByConnectionId(connectionId *string) error {
	sessions, err := repository.FindByConnectionId(connectionId)
	if err != nil {
		log.Printf("failed to get sessions by connectionId %s. error: %s", *connectionId, err.Error())
		return err
	}
	errorMessage := ""
	for _, v := range sessions {
		_, err = repository.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key:       repository.getKey(&v),
		})
		if err != nil {
			errorMessage += fmt.Sprintf("%s\n", errorMessage)
		}
	}
	if len(errorMessage) != 0 {
		log.Printf(errorMessage)
		return fmt.Errorf(errorMessage)
	}
	return nil
}

func (repository *SessionRepositoryDynamoDB) getKey(session *Session) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"connectionId": &types.AttributeValueMemberS{Value: *session.ConnectionId},
	}
}
