package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	Id        int64
	name      string
	kakaoAuth string
	appleAuth string
}

func (user User) GetKey() map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(user.Id)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{
		"id": id,
	}
}
