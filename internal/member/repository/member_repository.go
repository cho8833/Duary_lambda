package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cho8833/duary_lambda/internal/member/model"
	"strconv"
)

type MemberRepository interface {
	FindBySocialIdAndProvider(socialId int64, provider string) (*model.Member, error)
	SaveMember(member *model.Member) error
}

type MemberRepositoryDynamoDB struct {
	client dynamodb.Client
}

func NewMemberRepository(client *dynamodb.Client) *MemberRepositoryDynamoDB {
	return &MemberRepositoryDynamoDB{client: *client}
}

func (repo *MemberRepositoryDynamoDB) FindBySocialIdAndProvider(socialId int64, provider string) (*model.Member, error) {
	result, err := repo.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Member"),
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
	member := &model.Member{}
	err = attributevalue.UnmarshalMap(result.Item, member)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (repo *MemberRepositoryDynamoDB) SaveMember(member *model.Member) error {
	item, err := attributevalue.MarshalMap(member)
	if err != nil {
		return err
	}
	_, err = repo.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Member"),
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}
