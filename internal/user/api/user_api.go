package api

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/cho8833/CC-Calendar/internal/user/service"
	"strconv"
)

type UserAPI struct {
	svc service.UserService
}

func NewUserAPI(svc *service.UserService) *UserAPI {
	return &UserAPI{svc: *svc}
}

func (api UserAPI) GetUser(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userId, _ := strconv.ParseInt(event.QueryStringParameters["userId"], 10, 64)

	user, err := api.svc.GetUser(userId)

	var response events.APIGatewayProxyResponse

	if err != nil {
		response = events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("%#v", err),
		}
	} else {
		response = events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       fmt.Sprintf("%#v", user),
		}
	}

	return response, nil
}
