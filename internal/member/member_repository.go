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
		Key: map[string]types.AttributeValue{
			"socialId": &types.AttributeValueMemberN{Value: strconv.FormatInt(socialId, 10)},
			"provider": &types.AttributeValueMemberS{Value: provider},
		},
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
	update := expression.UpdateBuilder{}
	if member.Name != nil {
		update = update.Set(expression.Name("name"), expression.Value(*member.Name))
	}
	if member.AccessToken != nil {
		update = update.Set(expression.Name("accessToken"), expression.Value(member.AccessToken))
	}
	if member.Email != nil {
		update = update.Set(expression.Name("email"), expression.Value(member.Email))
	}
	if member.Gender != nil {
		update = update.Set(expression.Name("gender"), expression.Value(member.Gender))
	}
	if member.Birthday != nil {
		update = update.Set(expression.Name("birthday"), expression.Value(member.Birthday))
	}
	if member.FcmToken != nil {
		update = update.Set(expression.Name("fcmToken"), expression.Value(member.FcmToken))
	}

	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		log.Printf("failed to build update expression : %s", err.Error())
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

func (repo *RepositoryDynamoDB) getKey(socialId int64, provider string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"socialId": &types.AttributeValueMemberN{Value: strconv.FormatInt(socialId, 10)},
		"provider": &types.AttributeValueMemberS{Value: provider},
	}
}
