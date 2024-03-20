package api

import (
	"context"
	"fmt"
	"git-codecommit.ap-northeast-2.amazonaws.com/v1/repos/cc_calendar/internal/user/service"
	"github.com/aws/aws-lambda-go/events"
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
