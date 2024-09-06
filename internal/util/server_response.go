package util

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type ServerResponse struct {
	Message *string `json:"message"`
	Status  int     `json:"status"`
	Data    any     `json:"data"`
}

func ErrorResponse(message string, statusCode int) *ServerResponse {
	return &ServerResponse{
		&message,
		statusCode,
		nil,
	}
}

func ResponseFromError(error error, statusCode int) *ServerResponse {
	errorMessage := error.Error()
	return &ServerResponse{
		&errorMessage,
		statusCode,
		nil,
	}
}

func ResponseWithData(data any) *ServerResponse {
	okString := "OK"
	return &ServerResponse{
		&okString,
		200,
		&data,
	}
}

func SUCCESS() *ServerResponse {
	okString := "OK"

	return &ServerResponse{
		&okString,
		200,
		nil,
	}
}

func CreateLambdaResponse(res *ServerResponse) events.APIGatewayProxyResponse {
	b, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}
}
