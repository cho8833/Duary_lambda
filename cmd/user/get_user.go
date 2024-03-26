package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/CC-Calendar/config"
	"github.com/cho8833/CC-Calendar/internal/user/api"
	"github.com/cho8833/CC-Calendar/internal/user/repository"
	"github.com/cho8833/CC-Calendar/internal/user/repository/dynamodb"
	"github.com/cho8833/CC-Calendar/internal/user/service"
	"github.com/cho8833/CC-Calendar/internal/user/service/impl"
)

func main() {
	option, _ := config.GetDynamoDBOption()
	var userRepository repository.UserRepository = dynamodb.NewUserRepository(option.Client)
	var userService service.UserService = impl.NewUserService(&userRepository)
	userAPI := api.NewUserAPI(&userService)
	lambda.Start(userAPI.GetUser)
}
