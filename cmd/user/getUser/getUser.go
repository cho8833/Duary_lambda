package main

import (
	"git-codecommit.ap-northeast-2.amazonaws.com/v1/repos/cc_calendar/internal/user/api"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(api.GetUser)
}
