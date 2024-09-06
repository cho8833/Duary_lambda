package util

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type ServerResponse struct {
	Message *string `json:"message"`
	Status  int     `json:"status"`
	Data    any     `json:"data"`
	Error   bool    `json:"error"`
}

func ErrorResponse(message string, statusCode int) *ServerResponse {
	return &ServerResponse{
		&message,
		statusCode,
		nil,
		true,
	}
}

func ResponseFromError(error error, statusCode int) *ServerResponse {
	errorMessage := error.Error()
	return &ServerResponse{
		&errorMessage,
		statusCode,
		nil,
		true,
	}
}

func ResponseWithData(data any) *ServerResponse {
	okString := "OK"
	return &ServerResponse{
		&okString,
		200,
		&data,
		false,
	}
}

func SUCCESS() *ServerResponse {
	okString := "OK"

	return &ServerResponse{
		&okString,
		200,
		nil,
		false,
	}
}

func CreateLambdaResponse(res *ServerResponse) events.APIGatewayProxyResponse {
	b, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}
}
