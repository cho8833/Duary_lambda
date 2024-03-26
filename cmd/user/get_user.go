package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/CC-Calendar/config"
	"github.com/cho8833/CC-Calendar/internal/user/api"
	"github.com/cho8833/CC-Calendar/internal/user/repository/dynamodb"
	"github.com/cho8833/CC-Calendar/internal/user/service/impl"
)

func main() {
	option, _ := config.GetDynamoDBOption()
	userAPI := api.NewUserAPI(impl.NewUserService(dynamodb.NewUserRepository(option.Client)))
	lambda.Start(userAPI.GetUser)
}
