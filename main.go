package kakao_login_lambda

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/kakao_login_lambda/controller"
)

func main() {
	lambda.Start(controller.GetUser)
}
