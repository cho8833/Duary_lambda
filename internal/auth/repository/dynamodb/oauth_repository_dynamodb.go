package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cho8833/Duary/internal/auth/model"
	"strconv"
)

type OAuthTokenRepositoryDynamoDB struct {
	client dynamodb.Client
}

func NewOAuthTokenRepository(client *dynamodb.Client) *OAuthTokenRepositoryDynamoDB {
	return &OAuthTokenRepositoryDynamoDB{client: *client}
}

func (repo *OAuthTokenRepositoryDynamoDB) FindOAuthBySocialIdAndProvider(socialId int64, provider string) (*model.OAuthToken, error) {
	res, err := repo.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("OAuthToken"),
		Key: map[string]types.AttributeValue{
			"provider": &types.AttributeValueMemberS{Value: provider},
			"socialId": &types.AttributeValueMemberN{Value: strconv.FormatInt(socialId, 10)},
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Item == nil {
		return nil, &types.ResourceNotFoundException{
			Message: aws.String(fmt.Sprintf("resource not found for socialId: %d provider: %s", socialId, provider)),
		}
	}
	result := &model.OAuthToken{}
	err = attributevalue.UnmarshalMap(res.Item, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *OAuthTokenRepositoryDynamoDB) SaveOAuthToken(token *model.OAuthToken) error {
	item, err := attributevalue.MarshalMap(token)
	if err != nil {
		return err
	}

	_, err = repo.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("OAuthToken"),
		Item:      item,
	})
	if err != nil {
		return err
	}

	return nil
}
