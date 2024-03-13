package controller

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/cho8833/kakao_login_lambda/service"
	"strconv"
)

func GetUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userId, _ := strconv.ParseInt(event.QueryStringParameters["userId"], 10, 64)

	svc := service.NewUserService()
	user, _ := svc.GetUser(userId)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%#v", user),
	}

	return response, nil
}
