package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/CC-Calendar/internal/user/api"
)

func main() {
	lambda.Start(api.GetUser)
}
