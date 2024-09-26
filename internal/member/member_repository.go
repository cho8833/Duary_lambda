package member

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"strconv"
)

const tableName = "Member"

type Repository interface {
	FindBySocialIdAndProvider(socialId int64, provider string) (*Member, error)
	SaveMember(member *Member) (*Member, error)
	UpdateMember(member *UpdateMemberReq) (*Member, error)
	GetUpdateMemberTransaction(member *UpdateMemberReq) (*types.TransactWriteItem, error)
}

type RepositoryDynamoDB struct {
	client dynamodb.Client
}

func NewMemberRepository(client *dynamodb.Client) *RepositoryDynamoDB {
	return &RepositoryDynamoDB{client: *client}
}

func (repo *RepositoryDynamoDB) FindBySocialIdAndProvider(socialId int64, provider string) (*Member, error) {

	result, err := repo.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       repo.getKey(socialId, provider),
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, &types.ResourceNotFoundException{
			Message: aws.String(fmt.Sprintf("resource not found for id: %d", socialId)),
		}
	}
	member := &Member{}
	err = attributevalue.UnmarshalMap(result.Item, member)
	if err != nil {
		return nil, err
	}

	return member, nil
}
func (repo *RepositoryDynamoDB) SaveMember(member *Member) (*Member, error) {
	item, err := attributevalue.MarshalMap(member)
	if err != nil {
		return nil, err
	}
	_, err = repo.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (repo *RepositoryDynamoDB) UpdateMember(member *UpdateMemberReq) (*Member, error) {
	expr, err := repo.updateMemberExpression(member)
	if err != nil {
		return nil, err
	}

	response, err := repo.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       repo.getKey(member.SocialId, member.Provider),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		log.Printf("failed to update item. Member: %+v, Error: %s", member, err.Error())
		return nil, err
	}

	result := &Member{}
	_ = attributevalue.UnmarshalMap(response.Attributes, result)
	return result, nil
}

func (repo *RepositoryDynamoDB) updateMemberExpression(req *UpdateMemberReq) (*expression.Expression, error) {
	update := expression.UpdateBuilder{}
	if req.Name != nil {
		update = update.Set(expression.Name("name"), expression.Value(*req.Name))
	}
	if req.AccessToken != nil {
		update = update.Set(expression.Name("accessToken"), expression.Value(req.AccessToken))
	}
	if req.Email != nil {
		update = update.Set(expression.Name("email"), expression.Value(req.Email))
	}
	if req.Gender != nil {
		update = update.Set(expression.Name("gender"), expression.Value(req.Gender))
	}
	if req.Birthday != nil {
		update = update.Set(expression.Name("birthday"), expression.Value(req.Birthday))
	}
	if req.FcmToken != nil {
		update = update.Set(expression.Name("fcmToken"), expression.Value(req.FcmToken))
	}
	if req.CoupleId != nil {
		update = update.Set(expression.Name("coupleId"), expression.Value(req.CoupleId))
	}
	if req.Character != nil {
		update = update.Set(expression.Name("character"), expression.Value(req.Character))
	}

	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		log.Printf("failed to build update expression : %s", err.Error())
		return nil, err
	}
	return &expr, nil
}

func (repo *RepositoryDynamoDB) GetUpdateMemberTransaction(member *UpdateMemberReq) (*types.TransactWriteItem, error) {
	expr, err := repo.updateMemberExpression(member)
	if err != nil {
		return nil, err
	}
	result := &types.TransactWriteItem{Update: &types.Update{
		TableName:                 aws.String(tableName),
		Key:                       repo.getKey(member.SocialId, member.Provider),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		UpdateExpression:          expr.Update(),
	}}
	return result, err
}

func (repo *RepositoryDynamoDB) getKey(socialId int64, provider string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"socialId": &types.AttributeValueMemberN{Value: strconv.FormatInt(socialId, 10)},
		"provider": &types.AttributeValueMemberS{Value: provider},
	}
}
