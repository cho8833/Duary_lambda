package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/CC-Calendar/config"
	"github.com/cho8833/CC-Calendar/internal/user/api"
	"github.com/cho8833/CC-Calendar/internal/user/repository"
	"github.com/cho8833/CC-Calendar/internal/user/service"
)

func main() {
	userAPI := api.NewUserAPI(service.NewUserService(repository.NewUserRepository(config.DynamoDBOption{}.Client)))
	lambda.Start(userAPI.GetUser)
}
